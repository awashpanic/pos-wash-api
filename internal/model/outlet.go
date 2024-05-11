package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Outlet struct {
	OutletID  uuid.UUID      `json:"outlet_id" gorm:"primaryKey; default:gen_random_uuid()"`
	UserID    uuid.UUID      `json:"user_id"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	LogoID    *uuid.UUID     `json:"logo_id"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`

	User *User  `json:"user" gorm:"foreignKey:UserID; references:UserID"`
	Logo *Media `json:"logo" gorm:"foreignKey:LogoID; references:MediaID"`
}

func (Outlet) TableName() string {
	return "tr_outlet"
}
