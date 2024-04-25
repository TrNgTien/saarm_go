package modelRequests

import "github.com/google/uuid"

type NewUser struct {
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Status      string    `json:"status"`
	ApartmentID uuid.UUID `json:"apartmentID"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	ApartmentID uuid.UUID `json:"apartmentID"`
}
