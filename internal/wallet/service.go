package wallet

import (
	"fmt"

	"github.com/jesseinvent/go-payment-processor/internal/currency"
)

type WalletService interface {
	CreateWallet(currencyId, userId uint) (*Wallet, error)
	GetByID(id uint) (*Wallet, error)
	GetUserWallets(userId uint) ([]Wallet, error)
}
type walletService struct {
	walletStore WalletStore
	currencyStore currency.CurrencyStore
}

func NewWalletService(
		store WalletStore, 
		currencyStore currency.CurrencyStore,
	) WalletService {
	return &walletService{
		walletStore: store,
		currencyStore: currencyStore,
	}
}

func (s *walletService) CreateWallet(currencyId, userId uint) (*Wallet, error) {

	// Validate currency exists
	currency, err := s.currencyStore.GetByID(currencyId);

	if err != nil {
		return nil, fmt.Errorf("error getting currency - %w", err)
	}

	if currency == nil {
		return nil, fmt.Errorf("currency does not exist")
	}

	// Validate user does not have wallet with currency
	userCurrencyWallet, _ := s.walletStore.GetByUserIdAndCurrencyId(userId, currencyId)

	if userCurrencyWallet != nil {
		return nil, fmt.Errorf("userWallet with this currency already exists")
	}

	wallet := &Wallet{
		CurrencyId: currencyId,
		UserId: userId,
	}

	err = s.walletStore.Create(wallet)

	if err != nil {
		return nil, fmt.Errorf("error creating wallet - %w", err)
	}

	return wallet, nil
}

func (s *walletService) GetByID(id uint) (*Wallet, error) {
	return s.walletStore.GetByID(id)
}

func (s *walletService) GetUserWallets(userId uint) ([]Wallet, error) {
	return s.walletStore.FindByUserId(userId)
}