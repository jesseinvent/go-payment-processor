package payment

import (
	"fmt"
	"log"
	"time"

	"github.com/jesseinvent/go-payment-processor/internal/currency"
	ledgerEntry "github.com/jesseinvent/go-payment-processor/internal/ledger_entry"
	"github.com/jesseinvent/go-payment-processor/internal/transaction"
	"github.com/jesseinvent/go-payment-processor/internal/wallet"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FundWalletService interface {
	FundUserWallet(userId, walletId uint, amount float64) (*wallet.Wallet, error)
}
type fundWalletService struct {
	walletStore 		wallet.WalletStore
	currencyStore 		currency.CurrencyStore
	currencyService 	currency.CurrencyService
	transactionStore 	transaction.TransactionStore
	ledgerEntryStore 	ledgerEntry.LedgerEntryStore 
}

func NewFundWalletService(
		walletStore 		wallet.WalletStore, 
		currencyStore 		currency.CurrencyStore, 
		currencyService 	currency.CurrencyService, 
		transactionStore 	transaction.TransactionStore,
		ledgerEntryStore 	ledgerEntry.LedgerEntryStore, 
	) FundWalletService {
	return &fundWalletService{
		walletStore: 		walletStore,
		currencyStore: 		currencyStore,
		currencyService: 	currencyService,
		transactionStore: 	transactionStore,
		ledgerEntryStore: 	ledgerEntryStore,
	}
}

// Simulates funding a wallet by crediting the wallet and creating a transaction record
func (s *fundWalletService) FundUserWallet(userId, walletId uint, amount float64) (*wallet.Wallet, error) {

	var wallet *wallet.Wallet

	if amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than zero")
	}

	// Use transaction to ensure wallet balance is updated atomically with transaction record creation

	err := s.walletStore.WithTransaction(func(tx *gorm.DB) error {
			
		// Lock wallet record for update to prevent race conditions
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", walletId).First(&wallet).Error

		if err != nil {		
			return fmt.Errorf("error getting wallet - %w", err)
		}

		// Validate wallet belongs to user
		if wallet.UserId != userId {
			return fmt.Errorf("unauthorized to fund this wallet")
		}

		prevBalance := wallet.Balance

		amountInMinorUnit, err := s.currencyService.CalculateCurrencyAmountInBaseUnit(wallet.CurrencyId, float64(amount))

		log.Print("AMOUNT IN MINOR", amountInMinorUnit)

		if err != nil {
			return fmt.Errorf("error calculating amount in minor unit - %w", err)
		}

		// Credit wallet
		err = s.walletStore.Credit(tx, wallet.ID, uint(amountInMinorUnit))
		
		if err != nil {
			return fmt.Errorf("error updating wallet balance - %w", err)
		}

		ref := fmt.Sprintf("fund-%d-%d", walletId, time.Now().Unix())

		// Create transaction record
		transaction := &transaction.Transaction{
			UserId:                  userId,
			WalletId:                walletId,
			CurrencyId:              wallet.CurrencyId,
			Amount:                  int(amountInMinorUnit),
			Reference:               ref,
			TransactionType:         transaction.Received,
			Status:                  transaction.Completed,
		}

		err = tx.Create(transaction).Error

		if err != nil {
			return fmt.Errorf("error creating transaction record - %v", err)
		}	

		// Create ledger entry
		ledgerEntry := &ledgerEntry.LedgerEntry{
			UserId: userId,
			TransactionId: transaction.ID,
			WalletId: wallet.ID,
			CurrencyId: wallet.CurrencyId,
			EntryType: ledgerEntry.Credit,
			BalanceBefore: int(prevBalance),
			Amount: amountInMinorUnit,	
			BalanceAfter: int(prevBalance + uint(amountInMinorUnit)),
		}

		err = tx.Create(ledgerEntry).Error

		log.Print(ledgerEntry)

		if err != nil {
			return fmt.Errorf("error creating ledger entry record - %v", err)
		}	

		return nil
	})

	if err != nil {
		return nil, err
	}

	return wallet, nil
}