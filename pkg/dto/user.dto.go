package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserDtoData struct {
	ID          uuid.UUID
	Password    string
	Username    string
	LastLoginAt time.Time
}
