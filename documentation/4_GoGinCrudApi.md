# Go E-Commerce API - CRUD API Endpoints

## 4. Building API Endpoints – Implementing CRUD Routes

### 4.1 Setting Up API Routes

In Flask, you'd use Blueprints for modularization. In Gin, we achieve this by creating separate route files and grouping routes.

#### 4.1.1 Project Structure After This Step

```
go-ecommerce-api/
│── main.go
│── database/
│   ├── database.go
│── models/
│   ├── user.go
│   ├── product.go
│   ├── order.go
│── routes/
│   ├── user_routes.go
│   ├── product_routes.go
│   ├── order_routes.go
│── controllers/
│   ├── user_controller.go
│   ├── product_controller.go
│   ├── order_controller.go
│── .env
│── go.mod
│── go.sum
```

### 4.2 Modularizing Routes (Flask Blueprints Equivalent)

Flask Blueprints help in organizing routes; in Gin, we achieve the same using separate route files.

#### 4.2.1 Example: User Routes in Gin

📂 `routes/user_routes.go`

```go
package routes

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/controllers"
)

// SetupUserRoutes defines user-related routes
func SetupUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users") // Grouping under "/users"
	{
		userRoutes.POST("/", controllers.CreateUser) // Create User
		userRoutes.GET("/:id", controllers.GetUser)  // Get User by ID
		userRoutes.GET("/", controllers.GetUsers)    // Get All Users
		userRoutes.PUT("/:id", controllers.UpdateUser) // Update User
		userRoutes.DELETE("/:id", controllers.DeleteUser) // Delete User
	}
}
```

### 4.3 Implementing CRUD API Endpoints (User Example)

📂 `controllers/user_controller.go`

```go
package controllers

import (
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Function to hash passwords
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func CreateUser(c *gin.Context) {
	var User models.User

	// Bind JSON request to User struct
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()

	// Check if email already exists
	var existingUser models.User
	if err := db.Where("email = ?", User.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}

	// Hash password before saving
	hashedPassword, err := HashPassword(User.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	User.Password = hashedPassword

	// Create user
	if err := db.Create(&User).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": User})
}

// Other CRUD functions follow similar structure...
```

### 4.4 Integrating Routes in `main.go`

📂 `main.go`

```go
package main

import (
	"fmt"
	"go-ecommerce-api/database"
	"go-ecommerce-api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	database.ConnectDB()

	// Initialize router
	route := gin.Default()

	// Setup routes
	routes.SetupUserRoutes(route)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port", port)
	route.Run(":" + port)
}
```

### 4.5 API Documentation

#### 4.5.1 User Routes

| Method | Endpoint    | Description         |
| ------ | ----------- | ------------------- |
| POST   | /users/     | Create a new user   |
| GET    | /users/     | Get all users       |
| GET    | /users/\:id | Get user by ID      |
| PUT    | /users/\:id | Update user details |
| DELETE | /users/\:id | Delete a user       |

#### 4.5.2 Example Requests

**Create User**

```json
POST http://localhost:8080/users/
Content-Type: application/json

{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response**

```json
{
  "id": 1,
  "username": "john_doe",
  "email": "john@example.com",
  "created_at": "2025-04-02T14:00:00Z"
}
```

### 4.6 Authentication, Middleware & Best Practices

🔹 **JWT Authentication** – Secure endpoints using JWT tokens. 🔹 **Rate Limiting** – Prevent excessive requests to API. 🔹 **Validation Middleware** – Ensure correct data formats in requests. 🔹 **Logging Middleware** – Monitor API activity. 🔹 **CORS Middleware** – Allow cross-origin requests if needed.

🚀 **Next Steps** 🔹 Add authentication using JWT 🔹 Implement Product and Order routes 🔹 Add Unit Tests for API endpoints
