package ledgerentry

import "gorm.io/gorm"

type LedgerEntryType string

const (
	Credit LedgerEntryType = "credit"
	Debit LedgerEntryType = "debit"
)	
type LedgerEntry struct {
	gorm.Model
	UserId 			uint				`gorm:"not null"`
	WalletId 		uint  				`gorm:"not null;index"`
	TransactionId 	uint  				`gorm:"not null;index"`
	CurrencyId 		uint  				`gorm:"not null;index"`
	EntryType 		LedgerEntryType		`gorm:"not null"`
	BalanceBefore 	int					`gorm:"not null"`
	Amount			int 				`gorm:"not null"`
	BalanceAfter 	int					`gorm:"not null"`
}