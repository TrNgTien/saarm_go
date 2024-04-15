package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID   uuid.UUID 	`gorm:"type:uuid;"`
	Name string    `json:"name"`
	MonthlyPrice string    `json:"monthly_price"`
}

