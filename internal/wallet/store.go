package wallet

import (
	"errors"

	"gorm.io/gorm"
)
type WalletStore interface {
	Create(wallet *Wallet) error
	GetByID(id uint) (*Wallet, error)
	GetByUserIdAndCurrencyId(userId uint, currencyId uint) (*Wallet, error)
	FindByUserId(userId uint) ([]Wallet, error)
	WithTransaction(fn func(tx *gorm.DB) error) error
	Credit(tx *gorm.DB, walletId uint, amount uint) error
	Debit(tx *gorm.DB, walletId uint, amount uint) error
}

type walletStore struct {
	db *gorm.DB
}

func NewWalletStore(db *gorm.DB) WalletStore {
	return &walletStore{db: db}
}

func (s *walletStore) Create(wallet *Wallet) error {
	return s.db.Create(wallet).Error
}

func (s *walletStore) GetByID(id uint) (*Wallet, error) {

	var wallet Wallet
	
	err := s.db.First(&wallet, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &wallet, nil
}

func (s *walletStore) GetByUserIdAndCurrencyId(userId uint, currencyId uint) (*Wallet, error) {

	var wallet Wallet
	
	err := s.db.Where(&Wallet{
		UserId: userId,
		CurrencyId: currencyId,
	}).First(&wallet).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &wallet, nil
}


func (s *walletStore) FindByUserId(userId uint) ([]Wallet, error) {

	var wallets []Wallet

	err := s.db.Where(&Wallet{
		UserId: userId,
	}).Find(&wallets).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return wallets, nil
}

func (s *walletStore) WithTransaction(fn func(tx *gorm.DB) error) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

func (s *walletStore) Debit(tx *gorm.DB, walletId uint, amount uint) error {

	// If tx is nil, use s.db for the operation
	if tx == nil {
		tx = s.db
	}

	err := tx.Model(&Wallet{}).Where("id = ? AND balance >= ?", walletId, amount).UpdateColumn("balance", gorm.Expr("balance - ?", amount)).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("insufficient balance or balance changed")
		}
		return err
	}

	return nil
}

func (s *walletStore) Credit(tx *gorm.DB, walletId uint, amount uint) error {
	// If tx is nil, use s.db for the operation
	if tx == nil {
		tx = s.db
	}

	return tx.Model(&Wallet{}).Where("id = ?", walletId).UpdateColumn("balance", gorm.Expr("balance + ?", amount)).Error
}

