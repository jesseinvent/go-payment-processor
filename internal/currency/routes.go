package currency

import "github.com/gin-gonic/gin"

func RegisterCurrencyRoutes(rg *gin.RouterGroup, h CurrencyHandler) {
	currencyRoutes := rg.Group("/currencies")

	currencyRoutes.POST("/", h.CreateCurrency)
	currencyRoutes.GET("/", h.getCurrencies)
}