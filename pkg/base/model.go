package base

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
  ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:time"`
	ModifiedAt time.Time `json:"modified_at" gorm:"type:time"`
}
