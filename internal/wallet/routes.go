package wallet

import "github.com/gin-gonic/gin"

func RegisterWalletRoutes(rg *gin.RouterGroup, h *WalletHandler) {
	walletRoutes := rg.Group("/wallets")

	walletRoutes.POST("/", h.CreateWallet)
	walletRoutes.GET("/user/:userId", h.GetUserWallets)
}