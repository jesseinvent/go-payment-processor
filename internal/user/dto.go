package user

type CreateUserDto struct {
	Email 		string 	`json:"email"`
	PhoneNumber string	`json:"phoneNumber"`
	Name 		string	`json:"name"`
}