package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCurrency_Success(t *testing.T) {
	mockCurrencyStore := &mockCurrencyStore{
		CreateFunc: func(currency *Currency) error {
			return nil
		},
	}

	currencyService := &currencyService{
		store: mockCurrencyStore,
	}

	currency, err := currencyService.Create("USD", "USD", "http://example.com/usd.png", 100)

	assert.NoError(t, err)
	assert.NotNil(t, currency)
}

func TestCreateCurrency_Error(t *testing.T) {
	mockCurrencyStore := &mockCurrencyStore{
		CreateFunc: func(currency *Currency) error {
			return assert.AnError
		},
	}

	currencyService := &currencyService{
		store: mockCurrencyStore,
	}

	currency, err := currencyService.Create("USD", "USD", "http://example.com/usd.png", 100)

	assert.Error(t, err)
	assert.Nil(t, currency)
	assert.EqualError(t, err, assert.AnError.Error())
}

func TestGetCurrencyByID_Success(t *testing.T) {
	expectedCurrency := &Currency{
		Symbol: "USD",
		Name: "US Dollar",
		IconUrl: "http://example.com/usd.png",
		BaseUnitFactor: 100,
	}

	mockCurrencyStore := &mockCurrencyStore{
		GetByIDFunc: func(id uint) (*Currency, error) {
			return expectedCurrency, nil
		},
	}

	currencyService := &currencyService{
		store: mockCurrencyStore,
	}

	currency, err := currencyService.GetByID(1)

	assert.NoError(t, err)
		assert.NotNil(t, currency)
		assert.Equal(t, expectedCurrency.Symbol, currency.Symbol)
		assert.Equal(t, expectedCurrency.Name, currency.Name)
		assert.Equal(t, expectedCurrency.IconUrl, currency.IconUrl)
		assert.Equal(t, expectedCurrency.BaseUnitFactor, currency.BaseUnitFactor)
	}

func TestGetCurrencyByID_Error(t *testing.T) {
	mockCurrencyStore := &mockCurrencyStore{
		GetByIDFunc: func(id uint) (*Currency, error) {
			return nil, assert.AnError
		},
	}

	currencyService := &currencyService{
		store: mockCurrencyStore,
	}

	currency, err := currencyService.GetByID(1)

	assert.Error(t, err)
	assert.Nil(t, currency)
	assert.EqualError(t, err, assert.AnError.Error())
}