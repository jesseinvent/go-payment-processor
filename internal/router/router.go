package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-payment-processor/internal/currency"
	"github.com/jesseinvent/go-payment-processor/internal/transaction"
	"github.com/jesseinvent/go-payment-processor/internal/user"
	"github.com/jesseinvent/go-payment-processor/internal/wallet"
	"gorm.io/gorm"
)

/**
 * This file is responsible for registering all routes for the application.
 * It initializes the necessary services and handlers for each module and registers the routes with the Gin router.
 */
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	api := r.Group("/api/v1")

	// User
	userStore := user.NewUserStore(db)
	userService := user.NewUserService(userStore)
	userHandler := user.NewUserHandler(userService)
	user.RegisterUserRoutes(api, userHandler)

	// Currency
	currencyStore := currency.NewCurrencyStore(db)
	currencyService := currency.NewCurrencyService(currencyStore)
	currencyHandler := currency.NewCurrencyHandler(currencyService)
	currency.RegisterCurrencyRoutes(api, currencyHandler)

	// Wallet
	walletStore := wallet.NewWalletStore(db)
	walletService := wallet.NewWalletService(walletStore, &userStore, &currencyStore)
	walletHandler := wallet.NewWalletHandler(walletService)
	wallet.RegisterWalletRoutes(api, walletHandler)

	// Transaction
	transactionStore := transaction.NewTransactionStore(db)
	transactionService := transaction.NewTransactionService(transactionStore)
	transactionHandler := transaction.NewTransactionHandler(transactionService)
	transaction.RegisterTransactionRoutes(api, transactionHandler)
}