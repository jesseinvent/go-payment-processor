package payment

import (
	"context"
	"fmt"
	"time"

	"github.com/jesseinvent/go-payment-processor/internal/currency"
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
	walletStore wallet.WalletStore
	currencyStore currency.CurrencyStore
	currencyService currency.CurrencyService
	transactionStore transaction.TransactionStore
	redisService redis.RedisService
}

func NewInternalTransferService(
		walletStore wallet.WalletStore, 
		currencyStore currency.CurrencyStore, 
		currencyService currency.CurrencyService, 
		transactionStore transaction.TransactionStore, 
		redisService redis.RedisService,
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
		currencyId uint, 
		amount float64, 
		idempotencyKey string,
	) (string, error) {

	var senderCurrencyWallet, receiverCurrencyWallet *wallet.Wallet

	// Generate unique reference for transfer
	ref := utils.HashString(fmt.Sprintf("transfer-%d-%d-%d", senderUserId, receiverUserId, time.Now().Unix()))

	ctx := context.Background()

	existingRef, err := s.redisService.Get(ctx, idempotencyKey)

	// Check redis for existing transfer with same idempotency key to prevent duplicate processing
	 // If key exists, return existing reference without processing transfer again
	if err != nil {
		return "", fmt.Errorf("error checking idempotency key - %w", err)
	}

	// If key exists, return existing reference without processing transfer again
	if existingRef != "" {
		return existingRef, nil
	}

	// Using a database transaction for atomicity and to prevent race conditions
	err = s.walletStore.WithTransaction(func(tx *gorm.DB) error {

		if senderUserId == receiverUserId {
			return fmt.Errorf("you cannot transfer to yourself")
		}

		if amount <= 0 {
			return fmt.Errorf("amount must be greater than zero")
		}

		// Lock sender currency wallet for update
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id = ? AND currency_id = ?", senderUserId, currencyId).First(&senderCurrencyWallet).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("sender wallet not found for the specified currency")
			}

			return fmt.Errorf("error getting sender wallet - %w", err)
		}

		// Lock receiver currency wallet for update
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id = ? AND currency_id = ?", receiverUserId, currencyId).First(&receiverCurrencyWallet).Error

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

		senderCurrencyWalletBalance := senderCurrencyWallet.Balance - uint(amountInMinorUnit)

		// Debit sender wallet
		err = s.walletStore.Debit(senderCurrencyWallet.ID, uint(amountInMinorUnit))

		if err != nil {
			return fmt.Errorf("error debiting sender wallet - %w", err)
		}

		// Create debit transaction record for sender
		senderDebitTransaction := &transaction.Transaction{
			UserId:                  senderUserId,
			WalletId:                senderCurrencyWallet.ID,
			CurrencyId:              senderCurrencyWallet.CurrencyId,
			PreviousWalletBalance:   int(prevSenderBalance),
			Amount:                  int(amountInMinorUnit),
			CurrentWalletBalance:    int(senderCurrencyWalletBalance),
			Reference:               ref,
			TransactionType:         transaction.Debit,
			Status:                  transaction.Completed,
		}

		err = tx.Create(senderDebitTransaction).Error

		if err != nil {
			return fmt.Errorf("error creating sender transaction record - %w", err)
		}

		receiverCurrencyWalletBalance := receiverCurrencyWallet.Balance + uint(amountInMinorUnit)
		
		// Credit receiver wallet
		err = s.walletStore.Credit(receiverCurrencyWallet.ID, uint(amountInMinorUnit))

		if err != nil {
			return fmt.Errorf("error crediting receiver wallet - %w", err)
		}

		// Create credit transaction record for receiver
		receiverCreditTransaction := &transaction.Transaction{
			UserId:                  receiverUserId,
			WalletId:                receiverCurrencyWallet.ID,
			CurrencyId:              receiverCurrencyWallet.CurrencyId,
			PreviousWalletBalance:   int(prevReceiverBalance),
			Amount:                  int(amountInMinorUnit),
			CurrentWalletBalance:    int(receiverCurrencyWalletBalance),
			Reference:               ref,
			TransactionType:         transaction.Credit,
			Status:                  transaction.Completed,
		}

		err = tx.Create(receiverCreditTransaction).Error

		if err != nil {
			return fmt.Errorf("error creating receiver transaction record - %w", err)
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