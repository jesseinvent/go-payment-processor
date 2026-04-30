package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalletService_Create_Success(t *testing.T) {
	mockWalletStore := &MockWalletStore{
		CreateFunc: func(wallet *Wallet) error {
			return nil
		},
	}
	
	mockWalletService := &walletService{
		walletStore: mockWalletStore,
	}

	wallet, err := mockWalletService.CreateWallet(1, 1)

	assert.NoError(t, err)
	assert.NotNil(t, wallet)	
}