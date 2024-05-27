package common

import (
	"github.com/golang-jwt/jwt/v5"
)

type PaginationQuery struct {
	Limit  int
	Offset int
	Page   int
}

type UploadWaterMeter struct {
	CroppedFile  string `json:"croppedFile"`
	OriginalFile string `json:"originalFile"`
}

type JwtCustomClaims struct {
	Role   string `json:"role"`
	UserID string `json:"userID"`
	RoomID string `json:"roomID"`
	Exp    int64  `json:"exp"`
	Iat    int64  `json:"iat"`
	jwt.RegisteredClaims
}
