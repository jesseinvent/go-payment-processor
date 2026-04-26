package currency

import (
	"errors"

	"gorm.io/gorm"
)

type CurrencyStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) CurrencyStore {
	return CurrencyStore{db: db}
}

func (s *CurrencyStore) Create(currency *Currency) error {
	return s.db.Create(currency).Error
}

func (s *CurrencyStore) GetAll() ([]Currency, error) {

	var currencies []Currency
	
	err := s.db.Find(&currencies).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
		}
	}

	return currencies, nil
}