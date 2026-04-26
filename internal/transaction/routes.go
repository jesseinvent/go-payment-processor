package transaction

import "github.com/gin-gonic/gin"

func RegisterTransactionRoutes(rg *gin.RouterGroup, h TransactionHandler) {
	transactionRoutes := rg.Group("/transactions")

	transactionRoutes.GET("/user/:userId", h.GetUserTransactions)
}