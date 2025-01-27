package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the PostgreSQL connection URL
	db_url := os.Getenv("POSTGRES_URL")

	if db_url == "" {
		log.Fatal("POSTGRES_URL not set in environment")
	}

	fmt.Println("db url =>>>>>>>>>>", db_url)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db
	fmt.Println("Database connected successfully!")
}
