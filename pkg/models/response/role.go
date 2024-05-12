package modelResponse

import (
	"saarm/pkg/base"
	"time"
)

type RoleResponse struct {
	base.BaseModel
	LastLoginAt time.Time `json:"lastLoginAt"`
	Email       string    `json:"email"`
}
