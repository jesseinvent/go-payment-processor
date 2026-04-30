package currency

type mockCurrencyStore struct {
	CreateFunc func(*Currency) error
	GetAllFunc func() ([]Currency, error)
	GetByIDFunc func(id uint) (*Currency, error)
}

func (m *mockCurrencyStore) Create(currency *Currency) error {
	return m.CreateFunc(&Currency{
		Name:          currency.Name,
		Symbol:        currency.Symbol,
		IconUrl:       currency.IconUrl,
		BaseUnitFactor: currency.BaseUnitFactor,
	}) 	
}

func (m *mockCurrencyStore) GetByID(id uint) (*Currency, error) {
	return m.GetByIDFunc(id)
}	

func (m *mockCurrencyStore) GetAll() ([]Currency, error) {
	return m.GetAllFunc()
}
