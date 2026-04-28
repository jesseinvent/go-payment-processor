package payment

import (
	"fmt"
	"log"
	"time"

	"github.com/jesseinvent/go-payment-processor/internal/currency"
	"github.com/jesseinvent/go-payment-processor/internal/transaction"
	"github.com/jesseinvent/go-payment-processor/internal/wallet"
	"gorm.io/gorm/clause"
)
type FundWalletService struct {
	walletStore *wallet.WalletStore
	currencyStore *currency.CurrencyStore
	currencyService *currency.CurrencyService
	transactionStore *transaction.TransactionStore
}

func NewFundWalletService(walletStore *wallet.WalletStore, currencyStore *currency.CurrencyStore, currencyService *currency.CurrencyService, transactionStore *transaction.TransactionStore) *FundWalletService {
	return &FundWalletService{
		walletStore: walletStore,
		currencyStore: currencyStore,
		currencyService: currencyService,
		transactionStore: transactionStore,
	}
}

// Simulates funding a wallet by crediting the wallet and creating a transaction record
func (s *FundWalletService) FundUserWallet(userId, walletId uint, amount float64) (*wallet.Wallet, error) {

	var wallet *wallet.Wallet

	// Use transaction to ensure wallet balance is updated atomically with transaction record creation

	tx := s.walletStore.DB.Begin()

	log.Print("AMOUNT", amount)

	if amount <= 0 {
		tx.Rollback()
		return nil, fmt.Errorf("amount must be greater than zero")
	}
		
	// Lock wallet record for update to prevent race conditions
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", walletId).First(&wallet).Error

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error getting wallet - %w", err)
	}

	// Validate wallet belongs to user
	if wallet.UserId != userId {
		tx.Rollback()
		return nil, fmt.Errorf("unauthorized to fund this wallet")
	}

	prevBalance := wallet.Balance

	amountInMinorUnit, err := s.currencyService.CalculateCurrencyAmountInBaseUnit(wallet.CurrencyId, float64(amount))

	log.Print("AMOUNT IN MINOR", amountInMinorUnit)

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error calculating amount in minor unit - %w", err)
	}

	// Credit wallet
	err = tx.Model(&wallet).Update("balance", wallet.Balance + uint(amountInMinorUnit)).Error

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error updating wallet balance - %w", err)
	}

	ref := fmt.Sprintf("fund-%d-%d", walletId, time.Now().Unix())


	// Create transaction record
	transaction := &transaction.Transaction{
		UserId:                  userId,
		WalletId:                walletId,
		CurrencyId:              wallet.CurrencyId,
		PreviousWalletBalance:   int(prevBalance),
		Amount:                  int(amountInMinorUnit),
		CurrentWalletBalance:    int(prevBalance + uint(amountInMinorUnit)),
		Reference:               ref,
		TransactionType:         transaction.Credit,
		Status:                  transaction.Completed,
	}

	err = tx.Create(transaction).Error

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error creating transaction record - %w", err)
	}	

	err = tx.Commit().Error

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error committing transaction - %w", err)
	}

	return wallet, nil
}