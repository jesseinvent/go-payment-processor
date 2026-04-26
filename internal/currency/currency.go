package currency

import (
	"gorm.io/gorm"
)
type Currency struct {
	gorm.Model		
	name     string		`gorm:"not null"`
	symbol string 		`gorm:"not null"`
	iconUrl string		`gorm:"not null"`
	status bool			`gorm:"not null"`
}