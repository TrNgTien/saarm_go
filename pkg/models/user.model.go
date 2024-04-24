package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	base.BaseModel
	LastLoginAt time.Time `json:"last_login_at" gorm:"default:CURRENT_TIMESTAMP;type:time"`
	Email       string    `json:"email" gorm:"unique"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Status      string    `json:"status" gorm:"type:string;default:100_ACTIVATED"`
}

type UserResponse struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Users struct {
	Users []User `json:"users"`
}

func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}
