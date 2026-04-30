package mocks

import (
	"github.com/jesseinvent/go-payment-processor/internal/wallet"
	"gorm.io/gorm"
)

// Shared mock implementation of WalletStore for testing purposes. For testing the wallet service without relying on a real database connection.
type MockWalletStore struct {
	CreateFunc func(wallet *wallet.Wallet) error
	GetByIDFunc func(id uint) (*wallet.Wallet, error)
	GetByUserIdAndCurrencyIdFunc func(userId uint, currencyId uint) (*wallet.Wallet, error)
	FindByUserIdFunc func(userId uint) ([]wallet.Wallet, error)
	DebitFunc func(tx *gorm.DB, walletId uint, amount uint) error
	CreditFunc func(tx *gorm.DB, walletId uint, amount uint) error
	WithTransactionFunc func(fn func(*gorm.DB) error) error	
}

func (m *MockWalletStore) Create(wallet *wallet.Wallet) error {
	return m.CreateFunc(wallet)
}

func (m *MockWalletStore) GetByID(id uint) (*wallet.Wallet, error) {
	return m.GetByIDFunc(id)
}

func (m *MockWalletStore) GetByUserIdAndCurrencyId(userId uint, currencyId uint) (*wallet.Wallet, error) {
	return m.GetByUserIdAndCurrencyIdFunc(userId, currencyId)
}

func (m *MockWalletStore) FindByUserId(userId uint) ([]wallet.Wallet, error) {
	return m.FindByUserIdFunc(userId)
}

func (m *MockWalletStore) Debit(tx *gorm.DB, walletId uint, amount uint) error {
	return m.DebitFunc(tx, walletId, amount)
}

func (m *MockWalletStore) Credit(tx *gorm.DB, walletId uint, amount uint) error {
	return m.CreditFunc(tx, walletId, amount)
}

func (m *MockWalletStore) WithTransaction(fn func(*gorm.DB) error) error {
	return m.WithTransactionFunc(fn)
}
