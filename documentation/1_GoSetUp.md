# Go E-Commerce API

## Step 1: Initialize the Go Project

### 1.1 Install Go and Create the Project Directory

Ensure Go is installed:
```sh
go version
```
If not installed, download it from [golang.org](https://go.dev/dl/).

Create project directory:
```sh
mkdir go-ecommerce-api && cd go-ecommerce-api
```

### 1.2 Initialize a Go Module
```sh
go mod init go-ecommerce-api
```

### 1.3 Install Required Packages
```sh
# Gin - HTTP web framework
 go get -u github.com/gin-gonic/gin
```
```sh
# GORM - ORM (Object Relational Mapper) for Go
 go get -u gorm.io/gorm
```
```sh
# PostgreSQL driver for GORM
 go get -u gorm.io/driver/postgres
```
```sh
# godotenv - Load environment variables from a .env file
 go get -u github.com/joho/godotenv
```

### 1.4 Create the Main Entry File
```sh
touch main.go
```

Open `main.go` and add:
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"  // Web framework for building APIs
	"github.com/joho/godotenv"  // Loads environment variables from a .env file
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
```

### 1.5 Create an `.env` File
```sh
touch .env
```

Add:
```
PORT=8080  # Port number for the server
```

### 1.6 Run the Application
```sh
go run main.go
```

Visit `http://localhost:8080/` and expect:
```json
{"message": "Welcome to the E-Commerce API"}
```


---