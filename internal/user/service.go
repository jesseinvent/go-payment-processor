package user

type UserService struct {
	userStore UserStore
}

func NewUserService(userStore UserStore) UserService {
	return UserService{userStore: userStore}
}

func (s *UserService) Create(email, phoneNumber, name string) (*User, error) {
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

func (s *UserService) GetByID(id int) (*User, error) {
	return s.userStore.GetByID(id)
}