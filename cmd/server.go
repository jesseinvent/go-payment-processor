package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os/user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-payment-processor/internal/configs"
	"github.com/jesseinvent/go-payment-processor/internal/currency"
	"github.com/jesseinvent/go-payment-processor/internal/db"
	"github.com/jesseinvent/go-payment-processor/internal/pkg/redis"
	"github.com/jesseinvent/go-payment-processor/internal/router"
	"github.com/jesseinvent/go-payment-processor/internal/transaction"
	"github.com/jesseinvent/go-payment-processor/internal/wallet"
	"gorm.io/gorm"
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

	dbConn, err := ConnectToDBWithRetry(dsn)

	if err != nil {
		log.Fatal("Error connecting to database - ", err)
	}

	// Use migrations on production
	if config.ENVIRONMENT == "development" {
	 err = dbConn.AutoMigrate(
			&user.User{}, 
			&transaction.Transaction{}, 
			&wallet.Wallet{}, 
			&currency.Currency{},
		)

	 if err != nil {
		log.Fatal("Error running auto migrations - ", err)
	 }
	}

	// Connect to Redis
	redisService, err := redis.NewRedisService(config.REDIS_URL)

	if err != nil {
		log.Fatal(err)
	}

	r.Use(gin.Logger())

	r.Use(cors.Default())

	r.GET("/", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Payment Processor.",
		})
	})

	// Register all routes
	router.RegisterRoutes(r, dbConn, redisService)

	// Start server
	log.Println("Server running on: ", port)

	err = r.Run(":" + port)
	
	if err != nil {
		log.Fatal("Error starting server - ", err)
	}
}

func ConnectToDBWithRetry(dsn string) (*gorm.DB, error) { 
	var dbConn *gorm.DB
	var err error

	for i := 0; i < 5; i++ {
		dbConn, err = db.ConnectDB(dsn);
		
		if err == nil {
			return dbConn, nil
		}

		log.Printf("Failed to connect to database (attempt %d): %v", i+1, err)
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	return nil, fmt.Errorf("failed to connect to database after 5 attempts: %v", err)
}