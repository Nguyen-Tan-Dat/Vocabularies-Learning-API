package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection and stores it in the DB variable
func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=160500 dbname=cic port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB = db
	log.Println("Database connection established successfully.")
}

// GetDB returns the initialized database instance
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call ConnectDatabase first.")
	}
	return DB
}
