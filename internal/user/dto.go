package user

import "time"

type CreateUserDto struct {
	Email 		string 	`json:"email"`
	PhoneNumber string	`json:"phoneNumber"`
	Name 		string	`json:"name"`
}

type UserResponse struct {
	ID			uint 		`json:"id"`
	Email 		string 		`json:"email"`
	PhoneNumber string		`json:"phoneNumber"`
	Name 		string		`json:"name"`
	CreatedAt 	time.Time 	`json:"createdAt"`
}