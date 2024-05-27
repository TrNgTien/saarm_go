package modelResponse

import "github.com/google/uuid"

type AparmentResponse struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	LocationUrl   string    `json:"location"`
	Address       string    `json:"address"`
	TotalRoom     int16     `json:"totalRoom"`
	RoomAvailable int16     `json:"roomAvailable"`
}
