package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-payment-processor/internal/configs"
	"github.com/jesseinvent/go-payment-processor/internal/currency"
	"github.com/jesseinvent/go-payment-processor/internal/db"
	"github.com/jesseinvent/go-payment-processor/internal/router"
	"github.com/jesseinvent/go-payment-processor/internal/transaction"
	"github.com/jesseinvent/go-payment-processor/internal/wallet"
)

func RunServer() {

	fmt.Println("Hello world")

	config := configs.LoadConfigs();

	// ctx := context.Background()

	port := config.PORT

	if port == "" {
		port = "5001"
	}

	r := gin.Default()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DB_HOST, config.DB_USER, config.DB_PASSWORD,config.DB_NAME, config.DB_PORT)

	dbConn, err := db.ConnectDB(dsn);

	if err != nil {
		log.Fatal("Error connecting to database - ", err)
	}

	// Use migrations on production
	if config.ENVIRONMENT == "development" {
		dbConn.AutoMigrate(
			&user.User{}, 
			&transaction.Transaction{}, 
			&wallet.Wallet{}, 
			&currency.Currency{},
		)
	}

	r.Use(gin.Logger())

	r.Use(cors.Default())

	r.GET("/", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Payment Processor.",
		})
	})

	// Register all routes
	router.RegisterRoutes(r, dbConn)

	// Start server
	log.Println("Server running on: ", port)

	r.Run(":" + port)
}
