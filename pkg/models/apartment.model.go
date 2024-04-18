package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Apartment struct {
	base.BaseModel
	Name string `json:"name"`
}

func (m *Apartment) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}

type AparmentResponse struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	ModifiedAt  time.Time `json:"modifiedAt"`
	LastLoginAt time.Time `json:"lastLoginAt"`
	Email       string    `json:"email"`
}
