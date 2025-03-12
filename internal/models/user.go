package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username         string `gorm:"unique;not null"`
    Email            string `gorm:"unique;not null"`
    PasswordHash     string `gorm:"not null"`
    ProfilePictureURL string `gorm:"size:255"` // URL for profile picture
    Bio              string `gorm:"size:500"` // Biography/description
    PhoneNumber      string `gorm:"size:20"`  // Optional phone number
    Role             string `gorm:"default:user"`
}
