package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Account struct {
  base.AccountBase
}

func (m *Account) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}
