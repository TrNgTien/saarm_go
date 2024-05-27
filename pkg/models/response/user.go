package modelResponse

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserResponse struct {
		ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
		Email       string    `json:"email"`
		Status      string    `json:"status" gorm:"type:string;default:100_ACTIVATED"`
		ApartmentID uuid.UUID `json:"apartmentID"`
	}

	AuthResponse struct {
		Value       string    `json:"value"`
		Type        string    `json:"type"`
		LastLoginAt time.Time `json:"lastLoginAt"`
	}

	SignUpResponse struct {
		LastLoginAt time.Time `json:"lastLoginAt"`
		UserID      uuid.UUID `json:"userId"`
	}
)
