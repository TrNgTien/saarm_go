package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetaLink struct {
	CreatedAt time.Time  `json:"created_at"`
	ModifiedAt time.Time  `json:"modified_at"`
	ID   uuid.UUID 	`gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name"`
}

func (m *MetaLink) BeforeCreate(tx *gorm.DB) (err error) {
  m.ID = uuid.New()
  m.CreatedAt = time.Now()
  m.ModifiedAt = time.Now()
  return
}
