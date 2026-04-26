package wallet

import (
	"fmt"

	"github.com/jesseinvent/go-payment-processor/internal/currency"
	"github.com/jesseinvent/go-payment-processor/internal/user"
)
type WalletService struct {
	walletStore WalletStore
	userStore *user.UserStore
	currencyStore *currency.CurrencyStore
}

func NewWalletService(store WalletStore, userStore *user.UserStore, currencyStore *currency.CurrencyStore) WalletService {
	return WalletService{
		walletStore: store,
		userStore: userStore,
		currencyStore: currencyStore,
	}
}

func (s *WalletService) CreateWallet(currencyId, userId uint) (*Wallet, error) {

	// Validate currency exists
	currency, err := s.currencyStore.GetByID(currencyId);

	if err != nil {
		return nil, fmt.Errorf("Error getting currency - %w", err)
	}

	if currency == nil {
		return nil, fmt.Errorf("Currency does not exist.")
	}

	// Validate user does not have wallet with currency
	userCurrencyWallet, err := s.walletStore.GetByUserIdAndCurrencyId(userId, currencyId)

	if err != nil {
		return nil, fmt.Errorf("Could not get user currency wallet.")
	}

	if userCurrencyWallet != nil {
		return nil, fmt.Errorf("UserWallet with this currency already exists")
	}

	wallet := &Wallet{
		CurrencyId: currencyId,
		UserId: userId,
	}

	err = s.walletStore.Create(wallet)

	if err != nil {
		return nil, fmt.Errorf("Error creating wallet - %w", err)
	}

	return wallet, nil
}

func (s *WalletService) GetByID(id uint) (*Wallet, error) {
	return s.walletStore.GetByID(id)
}

func (s *WalletService) GetUserWallets(userId uint) ([]Wallet, error) {
	return s.walletStore.FindByUserId(userId)
}