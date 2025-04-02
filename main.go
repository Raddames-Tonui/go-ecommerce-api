package main

import (
	"fmt"
	"go-ecommerce-api/config"
	"go-ecommerce-api/database"
	"log"
	"os"

	"github.com/gin-gonic/gin" // Web framework for building APIs
	"github.com/joho/godotenv" // Loads environment variables from a .env file
)

func main() {
	// Load environment variables from .env file or use  config.LoadEnvVariables()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect to the database
	database.ConnectDatabase() 

	// Initialize Gin router
	r := gin.Default()

	// Define a simple test route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the E-Commerce API",
		})
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	fmt.Println("Server running on port", port)
	r.Run(":" + port) // Start the server
}