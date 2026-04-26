package currency
type CurrencyService struct {
	store CurrencyStore
}

func NewService(store CurrencyStore) CurrencyService {
	return CurrencyService{store: store}
}

func (s *CurrencyService) Create(name, symbol, iconUrl string) (*Currency, error) {
	currency := &Currency{
		Name: name,
		Symbol: symbol,
		IconUrl: iconUrl,
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