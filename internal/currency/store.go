package currency

import (
	"errors"

	"gorm.io/gorm"
)

type CurrencyStore interface {
	Create(*Currency) error
	GetAll() ([]Currency, error)
	GetByID(id uint) (*Currency, error)
}

type currencyStore struct {
	db *gorm.DB
}

func NewCurrencyStore(db *gorm.DB) CurrencyStore {
	return &currencyStore{db: db}
}

func (s *currencyStore) Create(currency *Currency) error {
	return s.db.Create(currency).Error
}

func (s *currencyStore) GetAll() ([]Currency, error) {

	var currencies []Currency
	
	err := s.db.Find(&currencies).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
		}
	}

	return currencies, nil
}

func (s *currencyStore) GetByID(id uint) (*Currency, error) {

	var currency Currency
	
	err := s.db.First(&currency, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &currency, nil
}