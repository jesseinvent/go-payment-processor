package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserService_Create_Success(t *testing.T) {
	mockUserStore := &MockUserStore{
		CreateFunc: func(user *User) error {
			return nil
		},
	}

 	userService := &userService{
		userStore: mockUserStore,
	}

	user, err := userService.Create("test@gmail.com", "09037138243", "John Doe")

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserService_Create_Error(t *testing.T) {
	mockUserStore := &MockUserStore{
		CreateFunc: func(user *User) error {
			return errors.New("db error")
		},
	}

	userService := &userService{userStore: mockUserStore}

	user, err := userService.Create("test@gmail.com", "09037138243", "John Doe")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.EqualError(t, err, "db error")
}

func TestUserService_GetByID_Success(t *testing.T) {
	createdUser := &User{
		Name: "John Doe",
		Email: "john@gmail.com",
		PhoneNumber: "09023487632",
	}

	mockUserStore := &MockUserStore{
		GetByIDFunc: func(id uint) (*User, error) {
			return createdUser, nil
		},
	}

	userService := &userService{userStore: mockUserStore}

	user, err := userService.GetByID(1)

	assert.NotNil(t, user)
	assert.NoError(t, err)
	assert.Equal(t, createdUser.Name, user.Name)
	assert.Equal(t, createdUser.Email, user.Email)
	assert.Equal(t, createdUser.PhoneNumber, user.PhoneNumber)
}

func TestUserService_GetByID_Error(t *testing.T) {
		mockUserStore := &MockUserStore{
		GetByIDFunc: func(id uint) (*User, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}

	userService := &userService{userStore: mockUserStore}

	user, err := userService.GetByID(1)

	assert.Nil(t, user)
	assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
}