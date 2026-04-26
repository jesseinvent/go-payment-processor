package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-payment-processor/internal/currency"
	"github.com/jesseinvent/go-payment-processor/internal/user"
	"gorm.io/gorm"
)


func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	api := r.Group("/api/v1")

	// User
	userStore := user.NewUserStore(db)
	userService := user.NewUserService(userStore)
	userHandler := user.NewUserHandler(userService)
	user.RegisterUserRoutes(api, userHandler)


	// Currency
	currencyStore := currency.NewStore(db)
	currencyService := currency.NewService(currencyStore)
	currencyHandler := currency.NewHandler(currencyService)
	currency.RegisterCurrencyRoutes(api, currencyHandler)
}