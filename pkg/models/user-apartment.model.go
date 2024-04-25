package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserApartment struct {
	base.BaseModel
	UserID      uuid.UUID `json:"userID"`
	ApartmentID uuid.UUID `json:"apartmentID"`
	User        User      `gorm:"foreignKey:UserID"`
	Apartment   Apartment `gorm:"foreignKey:ApartmentID"`
}

func (m *UserApartment) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}
