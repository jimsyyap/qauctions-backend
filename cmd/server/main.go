package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/jimsyyap/backend/internal/config"
    "github.com/jimsyyap/backend/internal/handlers"
    "github.com/jimsyyap/backend/internal/models"
)

func main() {
    // Load configuration
    cfg := config.LoadConfig()

    // Connect to the database
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Auto-migrate the schema
    db.AutoMigrate(&models.User{})

    // Initialize Gin router
    r := gin.Default()

    // Register routes
    userHandler := handlers.NewUserHandler(db)
    r.POST("/register", userHandler.Register)
    r.POST("/login", userHandler.Login)

    // Start the server
    r.Run(":8080")
}
