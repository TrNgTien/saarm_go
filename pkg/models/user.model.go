package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
  base.BaseModel
	Email       string    `json:"email"`
	LastLoginAt time.Time `json:"lastLoginAt" gorm:"default:CURRENT_TIMESTAMP;type:time"`
	Username    string    `json:"username" gorm:"unique"`
	Password    string    `json:"password"`
	Status      string    `json:"status" gorm:"type:string;default:100_ACTIVATED"`
}

func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	m.LastLoginAt = time.Now()
	return
}
