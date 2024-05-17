package helpers

import (
	"saarm/pkg/common"
	"saarm/pkg/utilities"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func getToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := utilities.GetValueEnv("APP_ENV_SECRET_KEY", "secretKey")
	return token.SignedString([]byte(secretKey))
}

func GenerateToken(ID uuid.UUID, roleName string) (string, error) {
	claims := jwt.MapClaims{
		"exp":  GetOneDay(),
		"iat":  GetCurrentTime(),
		"role": roleName,
	}

	if roleName == common.TENANT_ROLE {
		claims["roomId"] = ID
		return getToken(claims)
	}

	claims["userId"] = ID

	return getToken(claims)
}
