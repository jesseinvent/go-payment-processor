package mocks

import "github.com/jesseinvent/go-payment-processor/internal/user"


type MockUserStore struct {
	CreateFunc  func(user *user.User) error
	GetByIDFunc	func(id uint) (*user.User, error)
}

func (m *MockUserStore) Create(user *user.User) error {
	return m.CreateFunc(user) 	
}

func (m *MockUserStore) GetByID(id uint) (*user.User, error) {
	return m.GetByIDFunc(id)
}