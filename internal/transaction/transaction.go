package transaction

import "gorm.io/gorm"

type TransactionStatus string

const (
	Pending TransactionStatus = "pending"
	Completed TransactionStatus = "completed"
	Failed TransactionStatus = "failed"
)

type Transaction struct {
	gorm.Model
	userId      				int					
	currencyId int
	previousWalletBalance 		uint              `gorm:"type:not null"`
	amount                      int
	currentWalletBalance        int				   `gorm:"type:not null"`
	internal                    bool			   
	externalPaymentDestination  string
	status                      TransactionStatus  `gorm:"type:user_status default:'active'"`
}