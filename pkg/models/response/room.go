package modelResponse

import (
	"time"

	"github.com/google/uuid"
)

type (
	RoomResponse struct {
		ID            uuid.UUID `json:"id"`
		Name          string    `json:"roomName"`
		Status        string    `json:"roomStatus"`
		RoomPrice     string    `json:"roomPrice"`
		ApartmentName string    `json:"apartmentName"`
		Address       string    `json:"apartmentAddress"`
	}

	HistorySubmitResponse struct {
		ID           uuid.UUID `json:"id"`
		CreatedAt    time.Time `json:"createdAt"`
		WaterNumber  string    `json:"waterMeter"`
		WaterConsume int16     `json:"waterConsume"`
	}

	DuplicateRoomResponse struct {
		ID uuid.UUID `json:"id"`
	}

	BillByRoomResponse struct {
		ID                 uuid.UUID `json:"id"`
		CreatedAt          time.Time `json:"createdAt"`
		WaterConsume       string    `json:"waterConsume"`
		ElectricityConsume string    `json:"electricityConsume"`
		ExtraFee           string    `json:"extraFee"`
		RoomPrice          string    `json:"roomPrice"`
	}
)
