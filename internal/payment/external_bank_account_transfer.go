package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jesseinvent/go-payment-processor/internal/currency"
	"github.com/jesseinvent/go-payment-processor/internal/pkg/redis"
	"github.com/jesseinvent/go-payment-processor/internal/pkg/utils"
	thirdparty "github.com/jesseinvent/go-payment-processor/internal/third_party"
	"github.com/jesseinvent/go-payment-processor/internal/transaction"
	"github.com/jesseinvent/go-payment-processor/internal/wallet"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ExternalBankAccountTransferService struct {
	walletStore *wallet.WalletStore
	currencyStore *currency.CurrencyStore
	currencyService *currency.CurrencyService
	transactionStore *transaction.TransactionStore
	redisService *redis.RedisService
	thirdPartyService *thirdparty.ThirdPartyPaymentAPI
}

func NewExternalBankAccountTransferService(
		walletStore *wallet.WalletStore, 
		currencyStore *currency.CurrencyStore, 
		currencyService *currency.CurrencyService, 
		transactionStore *transaction.TransactionStore, 
		redisService *redis.RedisService,
		thirdPartyService *thirdparty.ThirdPartyPaymentAPI,
	) *ExternalBankAccountTransferService {
	return &ExternalBankAccountTransferService{
		walletStore: walletStore,
		currencyStore: currencyStore,
		currencyService: currencyService,
		transactionStore: transactionStore,
		redisService: redisService,
		thirdPartyService: thirdPartyService,
	}
}

type BeneficiaryDetails struct {
	BeneficiaryName string
	BeneficiaryAccountNumber string
	BeneficiaryBankCode string
	SwiftCode string
	SortCode string
}

// Simulates transfers between a wallet and an external bank account using transactions.
func (s *ExternalBankAccountTransferService) ProcessExternalBankAccountTransfer(userId uint, currencyId uint, amount float64,  idempotencyKey string, beneficiaryDetails BeneficiaryDetails) (string, error) {
	 
	ctx := context.Background()

	existingRef, err := s.redisService.Get(ctx, idempotencyKey)

	if err != nil {
		return "", fmt.Errorf("error checking idempotency key - %w", err)
	}

	if existingRef != "" {
		return existingRef, nil
	}
 
	// Generate unique reference for transfer
	ref := utils.HashString(fmt.Sprintf("bank-transfer-%d-%d-%d", userId, currencyId, time.Now().Unix()))

	currency, err := s.currencyStore.GetByID(currencyId)

	if err != nil {
		return "", fmt.Errorf("error getting currency - %w", err)
	}

	if currency == nil {
		return "", fmt.Errorf("currency not found")
	}

	// Using a database transaction for atomicity and to prevent race conditions
	err = s.walletStore.DB.Transaction(func(tx *gorm.DB) error {
		// Get sender's currency wallet and lock for update to prevent race conditions
		senderCurrencyWallet := &wallet.Wallet{}

		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id = ? AND currency_id = ?", userId, currencyId).First(&senderCurrencyWallet).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("currency wallet not found for user")
			}

			return fmt.Errorf("error getting currency wallet - %w", err)
		}

		amountInMinorUnit, err := s.currencyService.CalculateCurrencyAmountInBaseUnit(currencyId, amount)

		if err != nil {
			return fmt.Errorf("error calculating currency amount in base unit - %w", err)
		}

		// Check if sender has sufficient balance
		if senderCurrencyWallet.Balance < uint(amountInMinorUnit) {
			return fmt.Errorf("insufficient wallet balance")
		}

		prevWalletBalance := senderCurrencyWallet.Balance

		// Debit sender's wallet
		senderCurrencyWallet.Balance -= uint(amountInMinorUnit)

		err = tx.Save(senderCurrencyWallet).Error

		if err != nil {
			return fmt.Errorf("could not update wallet balance - %w", err)
		}

		// Create debit transaction record with pending status
		debitTransaction := &transaction.Transaction{
			UserId:                  userId,
			WalletId:                senderCurrencyWallet.ID,
			CurrencyId:              senderCurrencyWallet.CurrencyId,
			PreviousWalletBalance:   int(prevWalletBalance),
			Amount:                  int(amountInMinorUnit),
			CurrentWalletBalance:    int(senderCurrencyWallet.Balance),
			Reference:               ref,
			TransactionType:         transaction.Debit,
			Status:                  transaction.Pending,
		}

		err = tx.Save(debitTransaction).Error

		if err != nil {
			return fmt.Errorf("could not create transaction record - %w", err)
		}

		// Commit transaction to save changes to database
		return nil
	})

	if err != nil {
		return "", err
	}

	// In a real scalable production system, the processes below (initiating third party transfer and updating transaction status) would be handled asynchronously using a message queue to improve latency of the API and reliability of the transfer.

	thirdPartyApiRequest := thirdparty.SimulateBankTransferRequest{
		UserId: userId,
		Amount: amount,
		Reference: ref,
		Currency: currency.Symbol,
		BeneficiaryName: beneficiaryDetails.BeneficiaryName,
		BeneficiaryAccountNumber: beneficiaryDetails.BeneficiaryAccountNumber,
		BeneficiaryBankCode: beneficiaryDetails.BeneficiaryBankCode,
		SwiftCode: beneficiaryDetails.SwiftCode,
		SortCode: beneficiaryDetails.SortCode,
	}

	// Simulate call to third party payment API to initiate bank transfer
	thirdPartyResponse, err := s.thirdPartyService.SimulateBankTransfer(thirdPartyApiRequest)

	if err != nil {
		return "", fmt.Errorf("error initiating bank transfer with third party API - %w", err)
	}

	thirdPartyResponseJson, err := json.Marshal(thirdPartyResponse)

	if err != nil {
		return "", fmt.Errorf("error converting third party response to JSON - %w", err)
	}

	// If API Request fails then reverse wallet debit and update transaction status to failed (Improvement)

	// Convert beneficiary details to json string
	beneficiaryDetailsJson, err := json.Marshal(beneficiaryDetails)

	if err != nil {
		return "", fmt.Errorf("error converting beneficiary details to JSON - %w", err)
	}

	beneficiaryDetailsJsonStr := string(beneficiaryDetailsJson)

	// Update transaction record with transfer reference and status based on response from third party API
	err = s.transactionStore.DB.Model(&transaction.Transaction{}).Where("reference = ?", ref).Updates(map[string]interface{}{
		"transaction_beneficiary_details": beneficiaryDetailsJsonStr,
		"status": transaction.Completed,
		"metadata": string(thirdPartyResponseJson), // Store third party response for audit and reconciliation purposes
	}).Error

	if err != nil {
		return "", fmt.Errorf("error updating transaction record - %w", err)
	}

	// Store transfer reference in redis with idempotency key to prevent duplicate processing
	_, err = s.redisService.Set(ctx, idempotencyKey, ref, time.Hour * 24)

	if err != nil {
		return "", fmt.Errorf("error storing transfer reference in Redis - %w", err)
	}

	// Return transfer reference
	return ref, nil
}