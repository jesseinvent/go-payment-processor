package currency

import "time"

type CreateCurrencyDto struct {
	Name 			string 	`json:"name"`
	Symbol 			string	`json:"symbol"`
	IconUrl 		string	`json:"iconUrl"`
	BaseUnitFactor 	int		`json:"baseUnitFactor"`
}

type CurrencyResponse struct {
	ID			uint 		`json:"id"`
	Name 		string 		`json:"name"`
	Symbol 		string		`json:"symbol"`
	BaseUnitFactor 	int		`json:"baseUnitFactor"`
	IconUrl 	string		`json:"iconUrl"`
	CreatedAt	time.Time	`json:"createdAt"`
}