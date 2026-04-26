package transaction

type TransactionResponse struct {
	UserId	  				uint	`json:"userId"`				
	WalletId				uint	`json:"walletId"`
	CurrencyId 				uint	`json:"currencyId"`
	Amount                  int     `json:"amount"`
	TransactionType         string  `json:"transactionType"`
	Status					string  `json:"status"`
}