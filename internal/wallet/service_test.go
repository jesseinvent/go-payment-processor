package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalletService_CreateWallet_Success(t *testing.T) {
	mockWalletStore := &mockWalletStore{
		CreateFunc: func(wallet *Wallet) error {
			return nil
		},
	}

	// mockCurrencyStore := &mocks.MockCurrencyStore{
	// 	GetByIDFunc: func(id uint) (*currency.Currency, error) {
	// 		return &currency.Currency{}, nil
	// 	},
	// }
	
	mockWalletService := &walletService{
		walletStore: mockWalletStore,
		// currencyStore: mockCurrencyStore,
	}

	wallet, err := mockWalletService.CreateWallet(1, 1)

	assert.NoError(t, err)
	assert.NotNil(t, wallet)	
}

func TestWalletService_Create_Error(t *testing.T) {
	mockWalletStore := &mockWalletStore{
		CreateFunc: func(wallet *Wallet) error {
			return assert.AnError
		},
	}
	
	mockWalletService := &walletService{
		walletStore: mockWalletStore,
	}

	wallet, err := mockWalletService.CreateWallet(1, 1)

	assert.Error(t, err)
	assert.Nil(t, wallet)	
	assert.EqualError(t, err, assert.AnError.Error())
}

func TestWalletService_GetByID_Success(t *testing.T) {
	expectedWallet := &Wallet{
		UserId: 1,
		CurrencyId: 1,
		Balance: 1000,
	}

	mockWalletStore := &mockWalletStore{
		GetByIDFunc: func(id uint) (*Wallet, error) {
			return expectedWallet, nil
		},
	}
	
	mockWalletService := &walletService{
		walletStore: mockWalletStore,
	}

	wallet, err := mockWalletService.GetByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, wallet)	
	assert.Equal(t, expectedWallet, wallet)
}

func TestWalletService_GetByID_Error(t *testing.T) {
	mockWalletStore := &mockWalletStore{
		GetByIDFunc: func(id uint) (*Wallet, error) {
			return nil, assert.AnError
		},
	}
	
	mockWalletService := &walletService{
		walletStore: mockWalletStore,
	}

	wallet, err := mockWalletService.GetByID(1)

	assert.Error(t, err)
	assert.Nil(t, wallet)	
	assert.EqualError(t, err, assert.AnError.Error())
}

func TestWalletService_GetUserWallets_Success(t *testing.T) {
	expectedWallets := []Wallet{
		{ UserId: 1, CurrencyId: 1, Balance: 1000},
		{ UserId: 1, CurrencyId: 2, Balance: 2000},
	}

	mockWalletStore := &mockWalletStore{
		GetUserWalletsFunc: func(userId uint) ([]Wallet, error) {
			return expectedWallets, nil
		},
	}
	
	mockWalletService := &walletService{
		walletStore: mockWalletStore,
	}

	wallets, err := mockWalletService.GetUserWallets(1)

	assert.NoError(t, err)
	assert.NotNil(t, wallets)	
	assert.Equal(t, expectedWallets, wallets)
}

func TestWalletService_GetUserWallets_Error(t *testing.T) {
	mockWalletStore := &mockWalletStore{
		GetUserWalletsFunc: func(userId uint) ([]Wallet, error) {
			return nil, assert.AnError
		},
	}
	
	mockWalletService := &walletService{
		walletStore: mockWalletStore,
	}

	wallets, err := mockWalletService.GetUserWallets(1)

	assert.Error(t, err)
	assert.Nil(t, wallets)	
	assert.EqualError(t, err, assert.AnError.Error())
}	