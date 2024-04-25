package models

import (
	"saarm/pkg/base"

	"github.com/google/uuid"
)

type PermissionMapping struct {
	base.BaseModel
	PermissionID uuid.UUID  `json:"permissionID"`
	RoleID       uuid.UUID  `json:"roleID"`
	Permission   Permission `gorm:"foreignKey:PermissionID"`
	Role         Role       `gorm:"foreignKey:RoleID"`
}
