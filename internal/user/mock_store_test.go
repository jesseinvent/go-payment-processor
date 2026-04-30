package user

type MockUserStore struct {
	CreateFunc  func(user *User) error
	GetByIDFunc	func(id uint) (*User, error)
}

func (m *MockUserStore) Create(user *User) error {
	return m.CreateFunc(user) 	
}

func (m *MockUserStore) GetByID(id uint) (*User, error) {
	return m.GetByIDFunc(id)
}