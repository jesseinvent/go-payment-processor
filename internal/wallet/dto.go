package wallet

import "time"

type CreateWalletDto struct {
	UserId      uint	`json:"userId"`
	CurrencyId	uint	`json:"currencyId"`	
}

type WalletResponse struct {
	ID			uint		`json:"id"`
	UserId		uint		`json:"userId"`
	CurrencyId 	uint		`json:"currencyId"`
	Balance		int			`json:"balance"`
	CreatedAt	time.Time	`json:"createdAt"`
}
