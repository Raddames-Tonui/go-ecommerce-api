
## Step 2: Setting Up Database Connection

### 2.1 Create Database Connection File

Create a new file `database/db.go`:

```sh
mkdir database && touch database/db.go
```

Open `database/db.go` and add:
```go

package database

import (
	"fmt"
	"go-ecommerce-api/config"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes and connects to the database
func ConnectDatabase() {
	// Load environment variables
	config.LoadEnvVariables()

	// Build database connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connected successfully!")
	DB = db
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

```

### 2.2 Update `main.go` to Initialize the Database
Modify `main.go` to include the database connection:
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"  // Web framework for building APIs
	"github.com/joho/godotenv"  // Loads environment variables from a .env file
	"go-ecommerce-api/database" // Import database connection
)

func main() {
	// Load environment variables
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
```

Add to `.env`:
```
DB_HOST=localhost  # Database host (e.g., localhost for local development)
DB_USER=your_user  # Database username
DB_PASSWORD=your_password  # Database password
DB_NAME=ecommerce_db  # Database name
DB_PORT=5432  # PostgreSQL default port
```

### 2.3 Run the Application
Ensure PostgreSQL is running, then start the application:
```sh
go run main.go
```

Expected output:
```sh
Database connected successfully!
Server running on port 8080
```
---