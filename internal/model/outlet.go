package model

import (
	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type Outlet struct {
	OutletID  uuid.UUID             `json:"outlet_id" gorm:"primaryKey;default:gen_random_uuid()"`
	Name      string                `json:"name"`
	Address   string                `json:"address"`
	LogoID    *uuid.UUID            `json:"logo_id"`
	CreatedAt int                   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int                   `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"column:deleted_at"`

	Logo *Media `json:"logo" gorm:"foreignKey:LogoID;references:MediaID"`
}

func (Outlet) TableName() string {
	return "tr_outlet"
}
