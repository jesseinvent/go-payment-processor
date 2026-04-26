package currency

import "time"

type CreateCurrencyDto struct {
	Name 		string 	`json:"name"`
	Symbol 		string	`json:"symbol"`
	IconUrl 	string	`json:"iconUrl"`
}

type CurrencyResponse struct {
	ID			uint 		`json:"id"`
	Name 		string 		`json:"name"`
	Symbol 		string		`json:"symbol"`
	IconUrl 	string		`json:"iconUrl"`
	CreatedAt	time.Time	`json:"createdAt"`
}