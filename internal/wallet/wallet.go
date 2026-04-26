package wallet

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model		
	UserId     uint	`gorm:"not null"`
	CurrencyId uint `gorm:"not null"`
	Balance    uint	`gorm:"not null;default:0;check:balance >= 0"`
}