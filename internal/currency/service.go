package currency

import "fmt"

type CurrencyService interface {
	Create(name, symbol, iconUrl string, baseUnitFactor int) (*Currency, error) 
	GetAll() ([]Currency, error)
	GetByID(id uint) (*Currency, error)
	CalculateCurrencyAmountInBaseUnit(currencyId uint, amount float64) (int, error)
}
type currencyService struct {
	store CurrencyStore
}

func NewCurrencyService(store CurrencyStore) CurrencyService {
	return &currencyService{store: store}
}

func (s *currencyService) Create(name, symbol, iconUrl string, baseUnitFactor int) (*Currency, error) {
	currency := &Currency{
		Name: name,
		Symbol: symbol,
		IconUrl: iconUrl,
		BaseUnitFactor: baseUnitFactor,
	}

	err := s.store.Create(currency)

	if err != nil {
		return nil, err
	}

	return currency, nil
}

func (s *currencyService) GetAll() ([]Currency, error) {
	return s.store.GetAll();
}

func (s *currencyService) GetByID(id uint) (*Currency, error) {
	return s.store.GetByID(id);
}

func (s *currencyService) CalculateCurrencyAmountInBaseUnit(currencyId uint, amount float64) (int, error) {
	currency, err := s.store.GetByID(currencyId)
	
	if err != nil {
		return 0, err
	}

	if currency == nil {
		return 0, fmt.Errorf("currency not found")
	}

	return int(amount * float64(currency.BaseUnitFactor)), nil
}
