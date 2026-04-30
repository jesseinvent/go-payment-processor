package wallet

import "gorm.io/gorm"

// Local mock implementation of WalletStore for testing purposes. For testing the wallet service without relying on a real database connection.
type mockWalletStore struct {
	CreateFunc func(wallet *Wallet) error
	GetByIDFunc func(id uint) (*Wallet, error)
	GetByUserIdAndCurrencyIdFunc func(userId uint, currencyId uint) (*Wallet, error)
	FindByUserIdFunc func(userId uint) ([]Wallet, error)
	GetUserWalletsFunc func(userId uint) ([]Wallet, error)
	DebitFunc func(tx *gorm.DB, walletId uint, amount uint) error
	CreditFunc func(tx *gorm.DB, walletId uint, amount uint) error
	WithTransactionFunc func(fn func(*gorm.DB) error) error	
}

func (m *mockWalletStore) Create(wallet *Wallet) error {
	return m.CreateFunc(wallet)
}

func (m *mockWalletStore) GetByID(id uint) (*Wallet, error) {
	return m.GetByIDFunc(id)
}

func (m *mockWalletStore) GetByUserIdAndCurrencyId(userId uint, currencyId uint) (*Wallet, error) {
	return m.GetByUserIdAndCurrencyIdFunc(userId, currencyId)
}

func (m *mockWalletStore) GetUserWallets(userId uint) ([]Wallet, error) {
	return m.GetUserWalletsFunc(userId)
}

func (m *mockWalletStore) FindByUserId(userId uint) ([]Wallet, error) {
	return m.FindByUserIdFunc(userId)
}

func (m *mockWalletStore) Debit(tx *gorm.DB, walletId uint, amount uint) error {
	return m.DebitFunc(tx, walletId, amount)
}

func (m *mockWalletStore) Credit(tx *gorm.DB, walletId uint, amount uint) error {
	return m.CreditFunc(tx, walletId, amount)
}

func (m *mockWalletStore) WithTransaction(fn func(*gorm.DB) error) error {
	return m.WithTransactionFunc(fn)
}
