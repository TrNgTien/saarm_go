package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Set expiry time (1 day)
		"iat":    time.Now().Unix(),                     // Issued at
		"userID": userID,                                // User ID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("APP_ENV_SECRET_KEY")

	return token.SignedString([]byte(secretKey))
}
