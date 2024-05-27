package models

import (
	"saarm/pkg/base"
)

type Permission struct {
	base.BaseModel
	Name      string    `json:"Name"`
}
