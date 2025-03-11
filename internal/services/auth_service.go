package services

import (
    "crypto/rand"
    "encoding/base64"
    "errors"
    "time"

    jwt "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// VerifyPassword verifies the given password against the hash
func VerifyPassword(hash, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// GenerateJWT generates a JWT token
func GenerateJWT(userID uint, role string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    secret := []byte("your-secret-key") // Replace with a secure secret key
    return token.SignedString(secret)
}
