package models

import "gorm.io/gorm"

// Order represents an order placed by a user
type Order struct {
    gorm.Model
    UserID     uint    `json:"user_id" gorm:"not null"`
    User       User    `gorm:"foreignKey:UserID"` // Belongs to a User
    
    ProductID  uint    `json:"product_id" gorm:"not null"`
    Product    Product `gorm:"foreignKey:ProductID"` // Belongs to a Product

    Quantity   int     `json:"quantity" gorm:"not null"`
    TotalPrice float64 `json:"total_price" gorm:"not null"`
}