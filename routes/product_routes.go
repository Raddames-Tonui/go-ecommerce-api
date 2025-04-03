package routes

import (
	"go-ecommerce-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine) {
	productRoutes := router.Group("/products") // Grouping under "/products"
	{
		productRoutes.POST("/", controllers.CreateProduct) // Create Product
	}


}