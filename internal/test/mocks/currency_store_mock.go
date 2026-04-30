package mocks

import "github.com/jesseinvent/go-payment-processor/internal/currency"

type MockCurrencyStore struct {
	CreateFunc func(*currency.Currency) error
	GetAllFunc func() ([]currency.Currency, error)
	GetByIDFunc func(id uint) (*currency.Currency, error)
}

func (m *MockCurrencyStore) Create(c *currency.Currency) error {
	return m.CreateFunc(&currency.Currency{
		Name:          c.Name,
		Symbol:        c.Symbol,
		IconUrl:       c.IconUrl,
		BaseUnitFactor: c.BaseUnitFactor,
	}) 	
}

func (m *MockCurrencyStore) GetByID(id uint) (*currency.Currency, error) {
	return m.GetByIDFunc(id)
}	

func (m *MockCurrencyStore) GetAll() ([]currency.Currency, error) {
	return m.GetAllFunc()
}
