package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserData struct {
	ID          uuid.UUID
	Password    string
	Username    string
	LastLoginAt time.Time
}
