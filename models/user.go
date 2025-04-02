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

	Orders []Order `json:"orders" gorm:"foreignKey:UserID"` // One-to-many relationship with orders	
}