package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-payment-processor/internal/currency"
	"github.com/jesseinvent/go-payment-processor/internal/payment"
	"github.com/jesseinvent/go-payment-processor/internal/pkg/redis"
	thirdparty "github.com/jesseinvent/go-payment-processor/internal/third_party"
	"github.com/jesseinvent/go-payment-processor/internal/transaction"
	"github.com/jesseinvent/go-payment-processor/internal/user"
	"github.com/jesseinvent/go-payment-processor/internal/wallet"
	"gorm.io/gorm"
)

/**
 * This file is responsible for registering all routes for the application.
 * It initializes the necessary services and handlers for each module and registers the routes with the Gin router.
 */
func RegisterRoutes(r *gin.Engine, db *gorm.DB, redisService redis.RedisService) {

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

	// Transaction
	transactionStore := transaction.NewTransactionStore(db)
	transactionService := transaction.NewTransactionService(transactionStore)
	transactionHandler := transaction.NewTransactionHandler(transactionService)
	transaction.RegisterTransactionRoutes(api, transactionHandler)

	// Wallet
	walletStore := wallet.NewWalletStore(db)
	walletService := wallet.NewWalletService(walletStore, currencyStore)
	walletHandler := wallet.NewWalletHandler(walletService)
	wallet.RegisterWalletRoutes(api, walletHandler)

	// Payment
	fundWalletService := payment.NewFundWalletService(
		walletStore, 
		currencyStore,
		currencyService, 
		transactionStore,
	)
	internalTransferService := payment.NewInternalTransferService(
		walletStore, 
		currencyStore, 
		currencyService, 
		transactionStore, 
		redisService,
	)
	externalBankTransferService := payment.NewExternalBankAccountTransferService(
		walletStore,
		currencyStore,
		currencyService,
		transactionStore,
		redisService,
		thirdparty.NewThirdPartyPaymentAPI(),
	)

	paymentHandler := payment.NewPaymentHandler(fundWalletService, internalTransferService, externalBankTransferService)
	payment.RegisterPaymentRoutes(api, paymentHandler)
}