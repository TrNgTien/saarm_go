package modelRequests

import "github.com/google/uuid"

type (
	NewApartment struct {
		Name          string `json:"name"`
		LocationUrl   string `json:"locationUrl"`
		Address       string `json:"address"`
		TotalRoom     int16  `json:"totalRoom"`
		RoomAvailable int16  `json:"roomAvailable"`
	}

	LinkUser struct {
		UserID      uuid.UUID `json:"userId"`
		ApartmentID uuid.UUID `json:"apartmentId"`
	}
)
