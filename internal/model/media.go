package model

import (
	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type Media struct {
	MediaID   uuid.UUID             `json:"media_id" gorm:"primaryKey; default:gen_random_uuid()"`
	Name      string                `json:"name"`
	Size      int                   `json:"size"`
	Mimetype  string                `json:"mimetype"`
	MediaURL  string                `json:"media_url"`
	CreatedAt int                   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int                   `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

func (Media) TableName() string {
	return "tr_media"
}
