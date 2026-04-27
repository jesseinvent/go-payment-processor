package thirdparty

import (
	"fmt"
	"time"
)

type ThirdPartyPaymentAPI struct {}

func NewThirdPartyPaymentAPI() *ThirdPartyPaymentAPI {
	return &ThirdPartyPaymentAPI{}
}

type SimulateBankTransferRequest struct {
	UserId uint
	Amount float64
	Reference string
	Currency string
	BeneficiaryName string
	BeneficiaryAccountNumber string
	BeneficiaryBankCode string
	SwiftCode string
	SortCode string
}

type SimulateBankTransferResponse struct {
	TransferId string
	Status string
}

func (s *ThirdPartyPaymentAPI) SimulateBankTransfer(request SimulateBankTransferRequest) (SimulateBankTransferResponse, error) {

	// Simulates http call to third party payment API to initiate bank transfer
	time.Sleep(3 * time.Second)

	// Implementation might also include a webhook callback to update transfer status asynchronously after processing is complete, but for simplicity we will just simulate a successful transfer here.

	// Simulate successful transfer
	return SimulateBankTransferResponse{
		TransferId: fmt.Sprintf("thirdparty-transfer-%d", time.Now().Unix()),
		Status:     "success",
	}, nil
}
