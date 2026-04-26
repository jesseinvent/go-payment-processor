package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return UserStore{db: db}
}


func (s *UserStore) Create(user *User) error {
	return s.db.Create(user).Error
}

func (s *UserStore) GetByID(id int) (*User, error) {

	var user User
	
	err := s.db.First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
		}
	}

	return &user, nil
}