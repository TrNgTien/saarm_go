package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoomApartment struct {
	base.BaseModel
	RoomID      uuid.UUID `json:"roomID"`
	ApartmentID uuid.UUID `json:"apartmentID"`
	Room        Room      `gorm:"foreignKey:RoomID"`
	Apartment   Apartment `gorm:"foreignKey:ApartmentID"`
}

func (m *RoomApartment) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}
