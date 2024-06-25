package modelRequest

import "github.com/google/uuid"

type (
	NewApartment struct {
		Name          string    `json:"name"`
		LocationUrl   string    `json:"locationUrl"`
		Address       string    `json:"address"`
		TotalRoom     int16     `json:"totalRoom"`
		RoomAvailable int16     `json:"roomAvailable"`
		UserID        uuid.UUID `json:"userId"`
	}
)
