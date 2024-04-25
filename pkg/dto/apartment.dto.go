package dto

import (
	"github.com/google/uuid"
)

type ApartmentDto struct {
	ID            uuid.UUID `json:"id"`
	Name          string
	LocationUrl   string `json:"locationUrl"`
	Address       string `json:"address"`
	TotalRoom     int16  `json:"totalRoom"`
	RoomAvailable int16  `json:"roomAvailable"`
}
