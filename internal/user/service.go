package user

type UserService interface {
	Create(email, phoneNumber, name string) (*User, error)
	GetByID(id uint) (*User, error) 
}
type userService struct {
	userStore UserStore
}

func NewUserService(userStore UserStore) UserService {
	return &userService{userStore: userStore}
}

func (s *userService) Create(email, phoneNumber, name string) (*User, error) {
	user := &User{
		Email: email,
		PhoneNumber: phoneNumber,
		Name: name,
	}

	err := s.userStore.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetByID(id uint) (*User, error) {
	return s.userStore.GetByID(id)
}
