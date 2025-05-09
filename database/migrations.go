package database

import (
	"fmt"
	"log"
	"os"

	"go-ecommerce-api/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB2 *gorm.DB // Global database variable

// MigrateDb initializes the database and runs migrations
func MigrateDb() {
	// Load environment variables from .env file
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
	DB2 = db // Store the connection globally

	// Migrate the schema
	err = DB2.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	fmt.Println("Database migrated successfully!")
}

// GetDB returns the global database instance
func GetDataBase() *gorm.DB {
	return DB2
}
