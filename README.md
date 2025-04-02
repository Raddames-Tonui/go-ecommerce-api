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
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
# Step 3: Creating Models and Migrations

## In this step, we will:

- **Create Models**: Define Go structs that map to the database tables.
- **Migrations**: Automatically create or update tables in the database based on the models.


## 3.1: Define Models

Let's create a `models` folder and define the models for your E-commerce API. We'll create models for:

- **Product**: Represents the products in the store.
- **User**: Represents the users who are shopping.
- **Order**: Represents an order placed by a user.

### 3.1.1: Create the `models/product.go` file

```go
package models

import (
	"gorm.io/gorm"
)

// Product represents a product in the store
type Product struct {
	gorm.Model      // Embeds gorm.Model which includes ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null"`
	Stock       int     `json:"stock" gorm:"not null"`
}
```

### 3.1.2: Create the `models/user.go` file

```go
package models

import (
	"gorm.io/gorm"
)

// User represents a user of the store
type User struct {
	gorm.Model      // Embeds gorm.Model which includes ID, CreatedAt, UpdatedAt, DeletedAt
	Username  string `json:"username" gorm:"not null;unique"`
	Email     string `json:"email" gorm:"not null;unique"`
	Password  string `json:"password" gorm:"not null"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
```

### 3.1.3: Create the `models/order.go` file

```go
package models

import (
	"gorm.io/gorm"
)

// Order represents an order placed by a user
type Order struct {
	gorm.Model      // Embeds gorm.Model which includes ID, CreatedAt, UpdatedAt, DeletedAt
	UserID     uint   `json:"user_id" gorm:"not null"`
	ProductID  uint   `json:"product_id" gorm:"not null"`
	Quantity   int    `json:"quantity" gorm:"not null"`
	TotalPrice float64 `json:"total_price" gorm:"not null"`
}
```

---

## 3.2: Setup Migrations

Now, let's set up migrations that will automatically create or update the tables based on the defined models.

### 3.2.1: Create the `database/migrations.go` file

In the `database` package, create a `migrations.go` file. This file will handle the database connection and migration logic.

```go
package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"  // Loads environment variables from .env file
	"gorm.io/driver/postgres"   // PostgreSQL driver for GORM
	"gorm.io/gorm"              // GORM ORM for Go
	"github.com/your_project_name/models" // Import your models here
)

// ConnectDatabase will establish a connection to the PostgreSQL database and run migrations
func ConnectDatabase() *gorm.DB {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Build the PostgreSQL connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Connect to PostgreSQL using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Automate migration
	err = db.AutoMigrate(
		&models.Product{}, // Migrate Product model
		&models.User{},    // Migrate User model
		&models.Order{},   // Migrate Order model
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	fmt.Println("Database connected and migrated successfully!")
	return db
}
```

---

## 3.3: Initialize the Migration

### 3.3.1: Update `main.go` to Include Migration

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"  // Web framework for building APIs
	"github.com/joho/godotenv"  // Loads environment variables from .env file
	"gorm.io/gorm"              // GORM ORM for Go
	"your_project_name/database"  // Import the database package
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Database connection and migrate models
	db := database.ConnectDatabase()

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

---

## 3.4: Running Migrations

When you run your Go application, GORM will automatically create the necessary tables for Product, User, and Order in your PostgreSQL database.

### 3.4.1: Running the Application

To run the application and apply migrations, execute:

```bash
go run main.go
```

GORM will ensure the tables are created for the models you defined. If the tables already exist, GORM will attempt to migrate them (e.g., add missing columns, etc.).

---

## 3.5: Verifying the Database

Once the application is running, you can verify that the tables have been created in your PostgreSQL database by running a SQL query:

```sql
SELECT * FROM products;
SELECT * FROM users;
SELECT * FROM orders;
```

These tables should now exist, and you can start inserting data.





