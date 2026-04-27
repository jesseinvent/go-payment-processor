package currency

import (
	"gorm.io/gorm"
)
type Currency struct {
	gorm.Model		
	Name     			string	`gorm:"not null"`
	Symbol 	 			string  `gorm:"not null"`
	IconUrl  			string	`gorm:"not null"`
	BaseUnitFactor 		int  	`gorm:"not null"`
	Status 	 			bool	`gorm:"not null"`
}