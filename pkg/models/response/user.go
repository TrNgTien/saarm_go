package modelResponse

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserResponse struct {
		ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
		Email       string    `json:"email"`
		Name        string    `json:"name"`
		Status      string    `json:"status" gorm:"type:string;default:100_ACTIVATED"`
		LastLoginAt string    `json:"lastLoginAt"`
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
