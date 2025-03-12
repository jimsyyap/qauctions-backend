package main

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
    "github.com/jimsyyap/backend/internal/models"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found. Using system environment variables.")
    }

    // Retrieve database connection details from environment variables
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"))

    // Validate required environment variables
    if os.Getenv("DB_HOST") == "" || os.Getenv("DB_USER") == "" || os.Getenv("DB_PASSWORD") == "" || os.Getenv("DB_NAME") == "" || os.Getenv("DB_PORT") == "" {
        log.Fatal("Error: Missing required environment variables. Please check your .env file or system environment.")
    }

    // Connect to the database
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Auto-migrate the schema
    err = db.AutoMigrate(
        &models.User{},
        // Add other models here as needed (e.g., Listing, Bid, etc.)
    )
    if err != nil {
        log.Fatalf("Failed to run migrations: %v", err)
    }

    fmt.Println("Migrations completed successfully!")
}
