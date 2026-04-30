package user

import (
	"errors"

	"gorm.io/gorm"
)
type UserStore interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
}
type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return &userStore{db: db}
}

func (s *userStore) Create(user *User) error {
	return s.db.Create(user).Error
}

func (s *userStore) GetByID(id uint) (*User, error) {

	var user User
	
	err := s.db.First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
