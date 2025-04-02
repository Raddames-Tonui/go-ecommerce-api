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

	Orders      []Order  `gorm:"foreignKey:ProductID"` // One Product can be in multiple Orders
 
}
