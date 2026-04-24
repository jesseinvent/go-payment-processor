package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-payment-processor/configs"
)

func main() {

	fmt.Println("Hello world")

	config := configs.LoadConfigs();

	port := config.PORT

	if port == "" {
		port = "5001"
	}

	r := gin.New()

	r.Use(gin.Logger())

	r.GET("/", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	
	// Connect to database

	// Start server
	log.Println("Server running on: ", port)
	r.Run(":" + port)
}