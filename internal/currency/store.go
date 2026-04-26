package currency

import (
	"errors"

	"gorm.io/gorm"
)

type CurrencyStore struct {
	db *gorm.DB
}

func NewCurrencyStore(db *gorm.DB) CurrencyStore {
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

func (s *CurrencyStore) GetByID(id uint) (*Currency, error) {

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