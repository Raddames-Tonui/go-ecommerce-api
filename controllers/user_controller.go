package controllers

import (
	"go-ecommerce-api/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user); err
}