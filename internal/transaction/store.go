package transaction

import (
	"gorm.io/gorm"
)
type TransactionStore interface {
	Create(transaction *Transaction) error
	FindByWalletId(walletId uint) ([]Transaction, error)
	FindByUserId(userId uint) ([]Transaction, error)
	UpdateByReference(reference string, transaction *Transaction) (*Transaction, error)
}
type transactionStore struct {
	db *gorm.DB
}

func NewTransactionStore(db *gorm.DB) TransactionStore {
	return &transactionStore{db: db}
}

func (s *transactionStore) Create(transaction *Transaction) error {
	return s.db.Create(transaction).Error
}

func (s *transactionStore) FindByWalletId(walletId uint) ([]Transaction, error) {

	var transactions []Transaction
	
	err := s.db.Where(&Transaction{
		WalletId: walletId,
	}).Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *transactionStore) FindByUserId(userId uint) ([]Transaction, error) {

	var transactions []Transaction
	
	err := s.db.Where(&Transaction{
		UserId: userId,
	}).Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *transactionStore) UpdateByReference(reference string, fields *Transaction) (*Transaction, error) {

	var transaction Transaction

	err := s.db.Model(&Transaction{}).Where(&Transaction{Reference: reference}).Updates(fields).First(&transaction).Error

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}