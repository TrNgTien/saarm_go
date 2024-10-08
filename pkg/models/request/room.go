package modelRequest

import (
	"encoding/json"

	"github.com/google/uuid"
)

type (
	NewRoom struct {
		ID                    uuid.UUID       `json:"id"`
		Username              string          `json:"username" gorm:"unique"`
		Password              string          `json:"password"`
		Status                string          `json:"status" gorm:"type:string;default:100_ACTIVATED"`
		Name                  string          `json:"name"`
		RoomPrice             string          `json:"roomPrice"`
		MaxPeople             int8            `json:"maxPeople"`
		CurrentPeople         int8            `json:"currentPeople"`
		ApartmentID           uuid.UUID       `json:"apartmentId"`
		WaterNumberInit       string          `json:"waterNumberInit"`
		ElectricityNumberInit string          `json:"electricityNumberInit"`
		ExtraFee              json.RawMessage `json:"extraFee" gorm:"type:jsonb"`
	}

	DuplicateRoom struct {
		ID            uuid.UUID `json:"id"`
		Name          string    `json:"name"`
		RoomPrice     string    `json:"monthlyPrice"`
		MaxPeople     int8      `json:"maxPeople"`
		CurrentPeople int8      `json:"currentPeople"`
	}

	UpdateRoom struct {
		Username      string `json:"username" gorm:"unique"`
		Password      string `json:"password"`
		Status        string `json:"status" gorm:"type:string;default:100_ACTIVATED"`
		Name          string `json:"name"`
		RoomPrice     string `json:"monthlyPrice"`
		MaxPeople     int8   `json:"maxPeople"`
		CurrentPeople int8   `json:"currentPeople"`
	}

	SubmitWaterMeterNumber struct {
		WaterMeter string `json:"waterMeter"`
	}

	GetBillByRoom struct {
		MonthRequest string `json:"monthRequest"`
	}
)
