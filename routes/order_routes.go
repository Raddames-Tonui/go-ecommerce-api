package routes

import "github.com/gin-gonic/gin"


func SetupUserRoutes(router *gin.Engine){
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", )
	}
}
