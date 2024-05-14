package helpers

import (
	"saarm/pkg/utilities"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(userID uuid.UUID, roleName string) (string, error) {
	claims := jwt.MapClaims{
		"exp":    GetOneDay(),
		"iat":    GetCurrentTime(),
		"userID": userID,
		"role":   roleName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := utilities.GetValueEnv("APP_ENV_SECRET_KEY", "secretKey")

	return token.SignedString([]byte(secretKey))
}
