package wallet

import (
	"errors"

	"gorm.io/gorm"
)

type WalletStore struct {
	DB *gorm.DB
}

func NewWalletStore(db *gorm.DB) WalletStore {
	return WalletStore{DB: db}
}

func (s *WalletStore) Create(wallet *Wallet) error {
	return s.DB.Create(wallet).Error
}

func (s *WalletStore) GetByID(id uint) (*Wallet, error) {

	var wallet Wallet
	
	err := s.DB.First(&wallet, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &wallet, nil
}

func (s *WalletStore) GetByUserIdAndCurrencyId(userId uint, currencyId uint) (*Wallet, error) {

	var wallet Wallet
	
	err := s.DB.Where(&Wallet{
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


func (s *WalletStore) FindByUserId(userId uint) ([]Wallet, error) {

	var wallets []Wallet

	err := s.DB.Where(&Wallet{
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

