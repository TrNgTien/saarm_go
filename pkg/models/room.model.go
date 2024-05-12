package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	base.BaseModel
	LastLoginAt time.Time `json:"lastLoginAt" gorm:"default:CURRENT_TIMESTAMP;type:time"`
	Username    string    `json:"username" gorm:"unique"`
	Password    string    `json:"password"`
	Status      string    `json:"status" gorm:"type:string;default:100_ACTIVATED"`
	Name     string `json:"name"`
}

func (m *Room) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}
