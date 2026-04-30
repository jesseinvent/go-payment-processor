package wallet

type MockWalletService struct {
    CreateWalletFunc func(currencyId, userId uint) (*Wallet, error)
    GetByIDFunc      func(id uint) (*Wallet, error)
	GetUserWalletsFunc func(userId uint) ([]Wallet, error)
}

func (m *MockWalletService) CreateWallet(currencyId, userId uint) (*Wallet, error) {
    return m.CreateWalletFunc(currencyId, userId)
}

func (m *MockWalletService) GetByID(id uint) (*Wallet, error) {
    return m.GetByIDFunc(id)
}

func (m *MockWalletService) GetUserWallets(userId uint) ([]Wallet, error) {
	return m.GetUserWalletsFunc(userId)
}