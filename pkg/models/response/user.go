package modelReponses

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Email       string    `json:"email"`
	Status      string    `json:"status" gorm:"type:string;default:100_ACTIVATED"`
	ApartmentID uuid.UUID `json:"apartmentID"`
}

type AuthResponse struct {
	Value string `json:"value"`
	Type  string `json:"type"`
	LastLoginAt time.Time `json:"lastLoginAt"`
}
