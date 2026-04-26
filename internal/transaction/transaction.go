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
	Credit TransactionType = "credit"
	Debit TransactionType = "debit"
)	
type Transaction struct {
	gorm.Model
	UserId      				uint					
	WalletId					uint	
	CurrencyId 					uint
	PreviousWalletBalance 		int              `gorm:"type:not null"`
	Amount                      int
	CurrentWalletBalance        int				   `gorm:"type:not null"`
	Internal                    bool			   
	ExternalPaymentDestination  string
	Status                      TransactionStatus  `gorm:"type:user_status default:'active'"`
	TransactionType             TransactionType    `gorm:"type:transaction_type"`
}