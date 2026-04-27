package currency

import "fmt"

type CurrencyService struct {
	store CurrencyStore
}

func NewCurrencyService(store CurrencyStore) CurrencyService {
	return CurrencyService{store: store}
}

func (s *CurrencyService) Create(name, symbol, iconUrl string, baseUnitFactor int) (*Currency, error) {
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

func (s *CurrencyService) GetAll() ([]Currency, error) {
	return s.store.GetAll();
}

func (s *CurrencyService) GetByID(id uint) (*Currency, error) {
	return s.store.GetByID(id);
}

func (s *CurrencyService) CalculateCurrencyAmountInBaseUnit(currencyId uint, amount float64) (int, error) {
	currency, err := s.store.GetByID(currencyId)
	
	if err != nil {
		return 0, err
	}

	if currency == nil {
		return 0, fmt.Errorf("Currency not found")
	}

	return int(amount * float64(currency.BaseUnitFactor)), nil
}
