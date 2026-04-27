package transaction

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-payment-processor/internal/pkg/response"
)

type TransactionHandler struct {
	transactionService TransactionService
}

func NewTransactionHandler(transactionService TransactionService) TransactionHandler {
	return TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) GetUserTransactions(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid user id param"))
		return
	}

	transactions, err := h.transactionService.GetUserTransactions(uint(userId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("Failed to get user transactions"))
		return
	}

	var transactionResponses []TransactionResponse

	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, TransactionResponse{
			UserId: transaction.UserId,
			WalletId: transaction.WalletId,
			CurrencyId: transaction.CurrencyId,
			Amount: transaction.Amount,
			TransactionType: string(transaction.TransactionType),
			Status: string(transaction.Status),
			CreatedAt: transaction.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response.Success("User transactions retrieved successfully", transactionResponses))
}

