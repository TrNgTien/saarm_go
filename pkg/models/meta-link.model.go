package models

import (
	"saarm/pkg/base"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetaLink struct {
	base.BaseModel
	Name          string `json:"name"`
	Link          string `json:"link"`
	BucketName    string `json:"bucketName"`
	PrincipalType string `json:"principalType"`
	PrincipalID   string `json:"principalID"`
}

func (m *MetaLink) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
	m.ModifiedAt = time.Now()
	return
}
