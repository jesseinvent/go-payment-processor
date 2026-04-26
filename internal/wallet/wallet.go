package wallet

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model		
	userId     string		`gorm:"not null"`
	currencyId string 		`gorm:"not null"`
	balance    string		`gorm:"not null;default:0;check:balance >= 0"`
}