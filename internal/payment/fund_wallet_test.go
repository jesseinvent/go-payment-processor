package payment

import (
	"testing"

	"github.com/jesseinvent/go-payment-processor/internal/test/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestFundWalletService_FundUserWallet_Error(t *testing.T) {
	mockWalletStore := &mocks.MockWalletStore{
		WithTransactionFunc: func(fn func(tx *gorm.DB) error) error {
			return fn(nil)
		},
	}

	fundWalletService := &fundWalletService{
		walletStore: mockWalletStore,
	}

	wallet, err := fundWalletService.FundUserWallet(1, 1, -100)

	assert.Error(t, err)
	assert.Nil(t, wallet)
	assert.EqualError(t, err, "amount must be greater than zero")
}