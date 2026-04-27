package payment

import "github.com/gin-gonic/gin"

func RegisterPaymentRoutes(rg *gin.RouterGroup, h *PaymentHandler) {
	paymentRoutes := rg.Group("/payments")

	paymentRoutes.POST("/fund-wallet", h.FundWallet)
	paymentRoutes.POST("/internal-transfer", h.InternalTransfer)
	paymentRoutes.POST("/external-bank-transfer", h.ExternalBankAccountTransfer)
}