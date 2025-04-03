package controllers

import (
	"go-ecommerce-api/database"
	"go-ecommerce-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context){
	var Product models.Product

	// Bind JSON request to Product struct
	if err := c.ShouldBindJSON(&Product); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()

	// Check if product name already exists
	var existingProduct models.Product
	if err := db.Where("name = ?", Product.Name).First(&existingProduct).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product name already in use"})
		return
	}

	// Create product
	if err := db.Create(&Product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": Product})
}