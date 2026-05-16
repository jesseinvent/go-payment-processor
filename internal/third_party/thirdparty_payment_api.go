package thirdparty

import (
	"fmt"
	"log"
	"time"
)

type ThirdPartyPaymentAPI interface {
	SimulateBankTransfer(request SimulateBankTransferRequest) (SimulateBankTransferResponse, error) 
}

type thirdPartyPaymentAPI struct {}

func NewThirdPartyPaymentAPI() ThirdPartyPaymentAPI {
	return &thirdPartyPaymentAPI{}
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

func (s *thirdPartyPaymentAPI) SimulateBankTransfer(request SimulateBankTransferRequest) (SimulateBankTransferResponse, error) {

	log.Printf("Simulating bank transfer to third party API with request")

	// Simulates http call to third party payment API to initiate bank transfer
	time.Sleep(3 * time.Second)

	// Implementation might also include a webhook callback to update transfer status asynchronously after processing is complete, but for simplicity we will just simulate a successful transfer here.

	transferId := fmt.Sprintf("thirdparty-transfer-%d", time.Now().Unix())

	log.Printf("Simulated bank transfer successful with transfer ID %s", transferId)
	// Simulate successful transfer
	return SimulateBankTransferResponse{
		TransferId: transferId,
		Status:     "success",
	}, nil
}

