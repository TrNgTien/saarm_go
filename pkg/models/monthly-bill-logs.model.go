package models

import (
	"encoding/json"
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MonthlyBillLogs struct {
	base.BaseModel
	WaterNumber       string          `json:"waterNumber"`
	ElectricityNumber string          `json:"electricityNumber"`
	ExtraFee          json.RawMessage `json:"extraFee" gorm:"type:jsonb"`
	RoomID            uuid.UUID       `json:"roomID"`
	Room              Room            `gorm:"foreignKey:RoomID"`
}

func (m *MonthlyBillLogs) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}
