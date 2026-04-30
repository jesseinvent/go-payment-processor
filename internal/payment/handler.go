package payment

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-payment-processor/internal/pkg/response"
	"github.com/jesseinvent/go-payment-processor/internal/wallet"
)

type PaymentHandler struct {
	fundWalletService FundWalletService
	internalTransferService InternalTransferService
	externalBankTransferService ExternalBankAccountTransferService
}

func NewPaymentHandler(
	fundWalletService FundWalletService, 
	internalTransferService InternalTransferService, 
	externalBankTransferService ExternalBankAccountTransferService,
) *PaymentHandler {
	return &PaymentHandler{
		fundWalletService: fundWalletService, 
		internalTransferService: internalTransferService,
		externalBankTransferService: externalBankTransferService,
	}
}


func (h *PaymentHandler) FundWallet(c *gin.Context) {

	var fundWalletDto FundWalletDto

	err := c.ShouldBindJSON(&fundWalletDto)

	if err != nil {
		c.JSON(400, response.Error(fmt.Sprintf("Invalid request format - %s", err.Error())))
		return
	}

	userWallet, err := h.fundWalletService.FundUserWallet(fundWalletDto.WalletId, fundWalletDto.UserId, fundWalletDto.Amount)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(fmt.Sprintf("Error funding wallet - %s", err.Error())))
		return
	}

	walletResponse := wallet.WalletResponse{
		ID: userWallet.ID,
		CurrencyId: userWallet.CurrencyId,
		UserId: userWallet.UserId,
		Balance: int(userWallet.Balance),
		CreatedAt: userWallet.CreatedAt,
	}

	c.JSON(http.StatusOK, response.Success("Wallet successfully funded", walletResponse))	
}

func (h *PaymentHandler) InternalTransfer(c *gin.Context) {

	var transferDto InternalTransferDto

	idempotencyKey := c.GetHeader("Idempotency-Key")

	if idempotencyKey == "" {
		c.JSON(400, response.Error("Idempotency-Key header is required"))
		return
	}

	err := c.ShouldBindJSON(&transferDto)

	if err != nil {
		c.JSON(400, 
			response.Error(fmt.Sprintf("Invalid request format - %s", err.Error())),
		)
		return
	}

	transferRef, err := h.internalTransferService.ProcessInternalWalletTransfer(
		transferDto.SenderUserId, 
		transferDto.ReceiverUserId,
		transferDto.CurrencyId,
		transferDto.Amount,
		idempotencyKey,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, 
			response.Error(fmt.Sprintf("Error processing transfer - %s", err.Error())),
		)
		return
	}

	c.JSON(
		http.StatusOK, 
		response.Success("Transfer successful", 
		gin.H{"reference": transferRef}),
	)	
}	

func (h *PaymentHandler) ExternalBankAccountTransfer(c *gin.Context) {

	var bankTransferDto ExternalBankAccountTransferDto

	idempotencyKey := c.GetHeader("Idempotency-Key")

	if idempotencyKey == "" {
		c.JSON(400, response.Error("Idempotency-Key header is required"))
		return
	}

	err := c.ShouldBindJSON(&bankTransferDto)

	if err != nil {
		c.JSON(400, 
			response.Error(fmt.Sprintf("Invalid request format - %s", err.Error())),
		)
		return
	}

	beneficiary := BeneficiaryDetails{
		BeneficiaryName: bankTransferDto.BeneficiaryName,
		BeneficiaryAccountNumber: bankTransferDto.BeneficiaryAccountNumber,
		BeneficiaryBankCode: bankTransferDto.BeneficiaryBankCode,
		SwiftCode: bankTransferDto.SwiftCode,
		SortCode: bankTransferDto.SortCode,
	}

	transferRef, err := h.externalBankTransferService.ProcessExternalBankAccountTransfer(
		bankTransferDto.SenderUserId,
		bankTransferDto.CurrencyId,
		bankTransferDto.Amount,
		idempotencyKey,
		beneficiary,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.Error(fmt.Sprintf("Error processing external bank transfer - %s", err.Error())),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.Success("External bank transfer successful",
			gin.H{"reference": transferRef}),
	)
}