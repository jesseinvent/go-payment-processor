package wallet

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-payment-processor/internal/pkg/response"
)
type WalletHandler struct {
	service WalletService
}

func NewWalletHandler(service WalletService) *WalletHandler {
	return &WalletHandler{service: service}
}

func (h *WalletHandler) CreateWallet (c *gin.Context) {

	var createWalletDto CreateWalletDto

	err := c.ShouldBindJSON(&createWalletDto)

	if err != nil {
		c.JSON(400, response.Error(fmt.Sprintf("Invalid request format - %s", err.Error())))
		return
	}

	wallet, err := h.service.CreateWallet(createWalletDto.CurrencyId, createWalletDto.UserId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(fmt.Sprintf("Error creating wallet - %s", err.Error())))
		return
	}
	
	walletResponse := WalletResponse{
		ID: wallet.ID,
		CurrencyId: wallet.CurrencyId,
		UserId: wallet.UserId,
		Balance: int(wallet.Balance),
		CreatedAt: wallet.CreatedAt,
	}

	c.JSON(http.StatusOK, response.Success("Wallet successfully created", walletResponse))
}

func (h *WalletHandler) GetUserWallets (c *gin.Context) {

	userId := c.Param("userId");

	if userId == "" {
		c.JSON(http.StatusBadRequest, response.Error("Please provide user Id"))
		return
	}

	id, err := strconv.Atoi(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid userId Param"))
		return
	}

	wallets, err := h.service.GetUserWallets(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Could get user wallets"))
		return
	}

	var walletsResponse []WalletResponse

	for _, wallet := range wallets {
		walletsResponse = append(walletsResponse, WalletResponse{
			ID: wallet.ID,
			CurrencyId: wallet.CurrencyId,
			UserId: wallet.UserId,
			Balance: int(wallet.Balance),
			CreatedAt: wallet.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response.Success("User wallets retrieved", walletsResponse))
}

