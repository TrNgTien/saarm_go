package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionMapping struct {
	base.BaseModel
	PermissionID uuid.UUID  `json:"permissionID"`
	RoleID       uuid.UUID  `json:"roleID"`
	Permission   Permission `gorm:"foreignKey:PermissionID"`
	Role         Role       `gorm:"foreignKey:RoleID"`
}

func (m *PermissionMapping) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}
