package payment

import (
	"context"
	"fmt"
	"time"

	"github.com/jesseinvent/go-payment-processor/internal/currency"
	ledgerEntry "github.com/jesseinvent/go-payment-processor/internal/ledger_entry"
	"github.com/jesseinvent/go-payment-processor/internal/pkg/redis"
	"github.com/jesseinvent/go-payment-processor/internal/pkg/utils"
	"github.com/jesseinvent/go-payment-processor/internal/transaction"
	"github.com/jesseinvent/go-payment-processor/internal/wallet"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
type InternalTransferService interface {
	ProcessInternalWalletTransfer(
		senderUserId, 
		receiverUserId, 
		currencyId uint, 
		amount float64, 
		idempotencyKey string,
	) (string, error)
}
type internalTransferService struct {
	walletStore 		wallet.WalletStore
	currencyStore 		currency.CurrencyStore
	currencyService 	currency.CurrencyService
	transactionStore 	transaction.TransactionStore
	redisService 		redis.RedisService
}

func NewInternalTransferService(
		walletStore 		wallet.WalletStore, 
		currencyStore 		currency.CurrencyStore, 
		currencyService 	currency.CurrencyService, 
		transactionStore 	transaction.TransactionStore, 
		redisService 		redis.RedisService,
	) InternalTransferService {
	return &internalTransferService{
		walletStore: walletStore,
		currencyStore: currencyStore,
		currencyService: currencyService,
		transactionStore: transactionStore,
		redisService: redisService,
	}
}

// Simulates transfers between two wallets using transactions.
func (s *internalTransferService) ProcessInternalWalletTransfer(
		senderUserId, 
		receiverUserId, 
		currencyId 		uint, 
		amount 			float64, 
		idempotencyKey 	string,
	) (string, error) {

	var senderCurrencyWallet, receiverCurrencyWallet *wallet.Wallet

	// Generate unique reference for transfer
	ref := utils.HashString(fmt.Sprintf("transfer-%d-%d-%d", senderUserId, receiverUserId, time.Now().Unix()))

	ctx := context.Background()

	// Checks if request has already been processed for the given key. If not, set the key with the transfer reference to prevent duplicate processing.

	// if KeySet = true; Key was not found in redis and was created, safe for processing.
	// if KeySet = false; Key already exists in redis and was not created, return existing reference without processing transfer again.
	
	keySet, err := s.redisService.SetNX(ctx, idempotencyKey, "pending", 5*time.Minute) // 24hours

	if !keySet {
		// Key already exists, return existing reference without processing transfer again
		existingRefValue, err := s.redisService.Get(ctx, idempotencyKey)

		if err != nil {
			return "", fmt.Errorf("error retrieving existing transfer reference from Redis - %w", err)
		}

		if existingRefValue == "pending" {
			return "Transfer is still pending", nil 
		}

		return existingRefValue, nil
	}

	// Key did not exist and was set successfully, safe to process transfer

	// Using a database transaction for atomicity and to prevent race conditions
	err = s.walletStore.WithTransaction(func(tx *gorm.DB) error {

		if senderUserId == receiverUserId {
			return fmt.Errorf("you cannot transfer to yourself")
		}

		if amount <= 0 {
			return fmt.Errorf("amount must be greater than zero")
		}

		// Lock sender currency wallet for update
		err := tx.Clauses(
				clause.Locking{Strength: "UPDATE"},
			).Where(
				"user_id = ? AND currency_id = ?", senderUserId, currencyId,
			).First(&senderCurrencyWallet).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("sender wallet not found for the specified currency")
			}

			return fmt.Errorf("error getting sender wallet - %w", err)
		}

		// Lock receiver currency wallet for update
		err = tx.Clauses(
				clause.Locking{Strength: "UPDATE"},
			).Where(
				"user_id = ? AND currency_id = ?", receiverUserId, currencyId,
			).First(&receiverCurrencyWallet).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("receiver wallet not found for the specified currency")
			}

			return fmt.Errorf("error getting receiver wallet - %w", err)
		}

		prevSenderBalance := senderCurrencyWallet.Balance
		prevReceiverBalance := receiverCurrencyWallet.Balance

		// Calculate amount in minor unit
		amountInMinorUnit, err := s.currencyService.CalculateCurrencyAmountInBaseUnit(currencyId, float64(amount))

		if err != nil {
			return fmt.Errorf("error calculating amount in minor unit - %w", err)
		}

		// Validate sender has sufficient balance
		if senderCurrencyWallet.Balance < uint(amountInMinorUnit) {
			return fmt.Errorf("insufficient balance in sender wallet")
		}

		senderCurrentWalletBalance := senderCurrencyWallet.Balance - uint(amountInMinorUnit)

		// Debit sender wallet
		err = s.walletStore.Debit(tx, senderCurrencyWallet.ID, uint(amountInMinorUnit)) 

		if err != nil {
			return fmt.Errorf("error debiting sender wallet - %w", err)
		}

		// Create debit transaction record for sender
		senderDebitTransaction := &transaction.Transaction{
			UserId:                  senderUserId,
			WalletId:                senderCurrencyWallet.ID,
			CurrencyId:              senderCurrencyWallet.CurrencyId,
			Amount:                  int(amountInMinorUnit),
			Reference:               ref,
			TransactionType:         transaction.Sent,
			Status:                  transaction.Completed,
		}

		err = tx.Create(senderDebitTransaction).Error

		if err != nil {
			return fmt.Errorf("error creating sender transaction record - %w", err)
		}

		senderLedgerEntry := &ledgerEntry.LedgerEntry{
			UserId: senderUserId,
			TransactionId: senderDebitTransaction.ID,
			WalletId: senderCurrencyWallet.ID,
			CurrencyId: currencyId,
			EntryType: ledgerEntry.Debit,
			BalanceBefore: int(prevSenderBalance),
			Amount: amountInMinorUnit,	
			BalanceAfter: int(senderCurrentWalletBalance),
		}

		err = tx.Create(senderLedgerEntry).Error

		if err != nil {
			return fmt.Errorf("error creating sender ledger record - %w", err)
		}
		
		// Credit receiver wallet
		err = s.walletStore.Credit(tx, receiverCurrencyWallet.ID, uint(amountInMinorUnit)) // $9

		if err != nil {
			return fmt.Errorf("error crediting receiver wallet - %w", err)
		}

		receiverCurrencyWalletBalance := receiverCurrencyWallet.Balance + uint(amountInMinorUnit)

		// Create credit transaction record for receiver
		receiverCreditTransaction := &transaction.Transaction{
			UserId:                  receiverUserId,
			WalletId:                receiverCurrencyWallet.ID,
			CurrencyId:              receiverCurrencyWallet.CurrencyId,
			Amount:                  int(amountInMinorUnit), // receiveAmount $9
			Reference:               ref,
			TransactionType:         transaction.Received,
			Status:                  transaction.Completed,
		}

		err = tx.Create(receiverCreditTransaction).Error

		if err != nil {
			return fmt.Errorf("error creating receiver transaction record - %w", err)
		}

		receiverLedgerEntry := &ledgerEntry.LedgerEntry{
			UserId: receiverUserId,
			TransactionId: receiverCreditTransaction.ID,
			WalletId: receiverCurrencyWallet.ID,
			CurrencyId: currencyId,
			EntryType: ledgerEntry.Credit,
			BalanceBefore: int(prevReceiverBalance),
			Amount: amountInMinorUnit,	
			BalanceAfter: int(receiverCurrencyWalletBalance),
		}

		err = tx.Create(receiverLedgerEntry).Error

		if err != nil {
			return fmt.Errorf("error creating receiver ledger record - %w", err)
		}

		return nil
	})	

	if err != nil {
		return "", fmt.Errorf("error processing internal transfer - %w", err)
	}

	// Store transfer reference in redis with idempotency key to prevent duplicate processing
	err = s.redisService.Set(ctx, idempotencyKey, ref, time.Hour * 24)

	if err != nil {
		return "", fmt.Errorf("error setting idempotency key in redis - %w", err)
	}

	return ref, nil
}