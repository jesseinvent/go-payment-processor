package wallet

import "gorm.io/gorm"

type MockWalletStore struct {
	CreateFunc func(wallet *Wallet) error
	GetByIDFunc func(id uint) (*Wallet, error)
	GetByUserIdAndCurrencyIdFunc func(userId uint, currencyId uint) (*Wallet, error)
	FindByUserIdFunc func(userId uint) ([]Wallet, error)
	DebitFunc func(walletId uint, amount uint) error
	CreditFunc func(walletId uint, amount uint) error
	WithTransactionFunc func(fn func(*gorm.DB) error) error	
}

func (m *MockWalletStore) Create(wallet *Wallet) error {
	return m.CreateFunc(wallet)
}

func (m *MockWalletStore) GetByID(id uint) (*Wallet, error) {
	return m.GetByIDFunc(id)
}

func (m *MockWalletStore) GetByUserIdAndCurrencyId(userId uint, currencyId uint) (*Wallet, error) {
	return m.GetByUserIdAndCurrencyIdFunc(userId, currencyId)
}

func (m *MockWalletStore) FindByUserId(userId uint) ([]Wallet, error) {
	return m.FindByUserIdFunc(userId)
}

func (m *MockWalletStore) Debit(walletId uint, amount uint) error {
	return m.DebitFunc(walletId, amount)
}

func (m *MockWalletStore) Credit(walletId uint, amount uint) error {
	return m.CreditFunc(walletId, amount)
}

func (m *MockWalletStore) WithTransaction(fn func(*gorm.DB) error) error {
	return m.WithTransactionFunc(fn)
}
