package transaction

import "gorm.io/gorm"

type TransactionStatus string
type TransactionType string

const (
	Pending TransactionStatus = "pending"
	Completed TransactionStatus = "completed"
	Failed TransactionStatus = "failed"
)

const (
	Sent TransactionType = "sent"
	Received TransactionType = "received"
)	
type Transaction struct {
	gorm.Model
	UserId      					uint				`gorm:"not null"`	
	WalletId						uint				`gorm:"not null"`
	CurrencyId 						uint				`gorm:"not null"`
	Reference 						string				`gorm:"not null"`
	Amount                     		int					`gorm:"not null"`
	Internal                    	bool			   	`gorm:"not null;default:false"`
	TransactionBeneficiaryDetails   string				`gorm:"null"`
	Status                      	TransactionStatus  	`gorm:"not null;default:'pending'"`
	TransactionType             	TransactionType    	`gorm:"not null"`
	Metadata                    	string				`gorm:"type:string;null"`
}