package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	base.BaseModel
	UserID uuid.UUID `json:"userID"`
	RoleID uuid.UUID `json:"roleID"`
	User   User      `gorm:"foreignKey:UserID"`
	Role   Role      `gorm:"foreignKey:RoleID"`
}

func (m *UserRole) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}
