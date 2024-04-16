package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
  ID   uuid.UUID 	` json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	ModifiedAt time.Time  `json:"modified_at"`
	LastLoginAt time.Time  `json:"last_login_at"`
	Email string    `json:"email" gorm:"unique"`
	Password string  `json:"password"`
	Status string  `json:"status"`
}

type UserResponse struct {
  ID   uuid.UUID 	`json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	ModifiedAt time.Time  `json:"modifiedAt"`
	LastLoginAt time.Time  `json:"lastLoginAt"`
	Email string    `json:"email"`
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
