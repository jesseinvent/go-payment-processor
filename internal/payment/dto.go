package payment

type FundWalletDto struct {
	UserId		uint	`json:"userId"`
	WalletId	uint	`json:"walletId"`
	Amount		float64	`json:"amount"`
}

type InternalTransferDto struct {
	SenderUserId		uint	`json:"senderUserId"`
	ReceiverUserId		uint	`json:"receiverUserId"`
	CurrencyId			uint	`json:"currencyId"`
	Amount				float64	`json:"amount"`
}

type ExternalBankAccountTransferDto struct {
	SenderUserId				uint		`json:"senderUserId"`
	CurrencyId					uint		`json:"currencyId"`
	Amount						float64		`json:"amount"`
	BeneficiaryName 			string		`json:"beneficiaryName"`
	BeneficiaryAccountNumber 	string		`json:"beneficiaryAccountNumber"`
	BeneficiaryBankCode 		string		`json:"beneficiaryBankCode"`
	SwiftCode 					string		`json:"swiftCode"`
	SortCode 					string		`json:"sortCode"`
}