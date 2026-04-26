package user

import (
	"gorm.io/gorm"
)
type User struct {
	gorm.Model		
	Email       string		`gorm:"not null:uniqueIndex"`
	PhoneNumber string 		`gorm:"not null:uniqueIndex"`
	Name 	    string		`gorm:"not null"`
}

