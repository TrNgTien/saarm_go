package models

import (
	"saarm/pkg/base"

	"github.com/google/uuid"
)

type MonthlyBillMapping struct {
	base.BaseModel
	MonthlyBillID uuid.UUID `json:"monthlyBillIDID"`
	RoomID        uuid.UUID `json:"roomID"`
	Room        Room        `gorm:"foreignKey:RoomID"`
	MonthlyBill MonthlyBill `gorm:"foreignKey:MonthlyBillID"`
}
