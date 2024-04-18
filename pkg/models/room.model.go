package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	base.BaseModel
	Name         string `json:"name"`
	MonthlyPrice string `json:"monthly_price"`
}

func (m *Room) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}
