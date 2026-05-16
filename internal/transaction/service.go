package transaction

import "fmt"

type TransactionService interface {
	Create(
		userId uint, 
		walletId uint, 
		currencyId uint,
		amount int, 
		transactionType TransactionType,
	) (*Transaction, error)

	GetWalletTransactions(walletId uint) ([]Transaction, error)

	GetUserTransactions(userId uint) ([]Transaction, error)
}
type transactionService struct {
	transactionStore TransactionStore
}

func NewTransactionService(transactionStore TransactionStore) TransactionService {
	return &transactionService{transactionStore: transactionStore}
}

func (s *transactionService) Create(
		userId uint, 
		walletId uint, 
		currencyId uint,
		amount int, 
		transactionType TransactionType,
	) (*Transaction, error) {
	transaction := &Transaction{
		UserId:     		userId,
		WalletId:   		walletId,
		CurrencyId: 		currencyId,
		Amount:     		amount,
		TransactionType:    transactionType,
	}

	err := s.transactionStore.Create(transaction)

	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}

	return transaction, nil
}

func (s *transactionService) GetWalletTransactions(walletId uint) ([]Transaction, error) {
	transactions, err := s.transactionStore.FindByWalletId(walletId)
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet transactions: %w", err)
	}

	return transactions, nil
}

func (s *transactionService) GetUserTransactions(userId uint) ([]Transaction, error) {
	transactions, err := s.transactionStore.FindByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user transactions: %w", err)
	}
	return transactions, nil
}