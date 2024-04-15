package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uuid.UUID 	`gorm:"type:uuid;"`
	Email string    `json:"email" gorm:"unique"`
	Password string  `json:"password"`
	LastLoginAt time.Time  `json:"last_Login_At"`
}

type Users struct {
	Users []User `json:"users"`
}
