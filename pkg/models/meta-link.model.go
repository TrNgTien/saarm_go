package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetaLink struct {
	base.BaseModel
	Name string `json:"name"`
	RoomID        uuid.UUID `json:"roomId"`
	Room          Room      `gorm:"foreignKey:RoomID"`
}

func (m *MetaLink) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}
