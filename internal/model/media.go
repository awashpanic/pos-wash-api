package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Media struct {
	MediaID   uuid.UUID      `json:"media_id" gorm:"primaryKey; default:gen_random_uuid()"`
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Size      int            `json:"size"`
	Mimetype  string         `json:"mimetype"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`

	// json field
	MediaURL string `json:"media_url" gorm:"-"`
}

func (Media) TableName() string {
	return "tr_media"
}
