package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Apartment struct {
	base.BaseModel
	Name          string `json:"name"`
	LocationUrl   string `json:"locationUrl"`
	Address       string `json:"address"`
	TotalRoom     int16  `json:"totalRoom"`
	RoomAvailable int16  `json:"roomAvailable"`
	Status        string `json:"status" gorm:"type:string;default:100_ACTIVATED"`
}

func (m *Apartment) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}

