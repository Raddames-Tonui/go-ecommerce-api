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
		// userRoutes.GET("/:id", controllers.GetUser)  // Get User by ID
		// userRoutes.GET("/", controllers.GetUsers)    // Get All Users
		// userRoutes.PUT("/:id", controllers.UpdateUser) // Update User
		// userRoutes.DELETE("/:id", controllers.DeleteUser) // Delete User
	}
}