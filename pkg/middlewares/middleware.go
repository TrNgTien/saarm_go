package middlewares

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	// if err := token.Claims.verifyToken; err != nil {
	//     return nil, err
	// }

	return claims, nil
}
