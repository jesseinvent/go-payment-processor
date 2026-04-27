package currency

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-payment-processor/internal/pkg/response"
)

type CurrencyHandler struct {
	currencyService CurrencyService
}

func NewCurrencyHandler(currencyService CurrencyService) CurrencyHandler {
	return CurrencyHandler{currencyService: currencyService}
}

func (h *CurrencyHandler) CreateCurrency(c *gin.Context) {

	var createCurrencyDto CreateCurrencyDto

	err := c.ShouldBindJSON(&createCurrencyDto)

	if err != nil {
		c.JSON(400, response.Error(fmt.Sprintf("Invalid request format - %s", err.Error())))
		return
	}

	currency, err := h.currencyService.Create(createCurrencyDto.Name, createCurrencyDto.Symbol, createCurrencyDto.IconUrl, createCurrencyDto.BaseUnitFactor)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Could not create currency."))
		return
  	}

	currencyResponse := CurrencyResponse{
		ID: currency.ID,
		Name: currency.Name,
		Symbol: currency.Symbol,
		IconUrl: currency.IconUrl,
		BaseUnitFactor: currency.BaseUnitFactor,
		CreatedAt: currency.CreatedAt,
	}

	c.JSON(http.StatusCreated, response.Success("Currency successfully created", currencyResponse))
}

func (h *CurrencyHandler) getCurrencies(c *gin.Context) {
	currencies, err := h.currencyService.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Could not retrieved currencies."))
		return
	}

	var	currencyResponses []CurrencyResponse

	for _, currency := range currencies {
		currencyResponses = append(currencyResponses, CurrencyResponse{
			ID: currency.ID,
			Name: currency.Name,
			Symbol: currency.Symbol,
			IconUrl: currency.IconUrl,
			BaseUnitFactor: currency.BaseUnitFactor,
			CreatedAt: currency.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response.Success("Successfully retrieved currencies", currencyResponses))

}