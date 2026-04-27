package transaction

import "gorm.io/gorm"

type TransactionStore struct {
	DB *gorm.DB
}

func NewTransactionStore(db *gorm.DB) TransactionStore {
	return TransactionStore{DB: db}
}

func (s *TransactionStore) Create(transaction *Transaction) error {
	return s.DB.Create(transaction).Error
}

func (s *TransactionStore) FindByWalletId(walletId uint) ([]Transaction, error) {

	var transactions []Transaction
	
	err := s.DB.Where(&Transaction{
		WalletId: walletId,
	}).Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *TransactionStore) FindByUserId(userId uint) ([]Transaction, error) {

	var transactions []Transaction
	
	err := s.DB.Where(&Transaction{
		UserId: userId,
	}).Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}