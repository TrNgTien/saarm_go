package modelResponse

import "github.com/google/uuid"

type (
	RoomResponse struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}
	DuplicateRoomResponse struct {
		ID uuid.UUID `json:"id"`
	}
)
