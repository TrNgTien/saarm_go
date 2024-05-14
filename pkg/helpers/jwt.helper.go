package helpers

import (
	// "os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"exp":    GetOneDay(),
		"iat":    GetCurrentTime(),
		"userID": userID,
		"role":   "guest",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// secretKey := os.Getenv("APP_ENV_SECRET_KEY")

	return token.SignedString([]byte("secretKey"))
}
