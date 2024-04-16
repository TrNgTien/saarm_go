package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	CreatedAt time.Time  `json:"created_at"`
	ModifiedAt time.Time  `json:"modified_at"`
	ID   uuid.UUID 	`gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name"`
	MonthlyPrice string    `json:"monthly_price"`
}

func (m *Room) BeforeCreate(tx *gorm.DB) (err error) {
  m.ID = uuid.New()
  m.CreatedAt = time.Now()
  m.ModifiedAt = time.Now()
  return
}
