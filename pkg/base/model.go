package base

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt  time.Time `json:"createdAt" gorm:"type:time"`
	ModifiedAt time.Time `json:"modifiedAt" gorm:"type:time"`
}

type AccountBase struct {
	BaseModel
	LastLoginAt time.Time `json:"lastLoginAt" gorm:"default:CURRENT_TIMESTAMP;type:time"`
	Username    string    `json:"username" gorm:"unique"`
	Password    string    `json:"password"`
	Status      string    `json:"status" gorm:"type:string;default:100_ACTIVATED"`
}
