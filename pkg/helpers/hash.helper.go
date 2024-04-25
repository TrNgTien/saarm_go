package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashed)
}

func ValidatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
