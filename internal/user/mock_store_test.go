package user

type mockUserStore struct {
	CreateFunc  func(user *User) error
	GetByIDFunc	func(id uint) (*User, error)
}

func (m *mockUserStore) Create(user *User) error {
	return m.CreateFunc(user) 	
}

func (m *mockUserStore) GetByID(id uint) (*User, error) {
	return m.GetByIDFunc(id)
}