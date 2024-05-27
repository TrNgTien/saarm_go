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
