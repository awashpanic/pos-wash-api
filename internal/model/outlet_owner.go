package model

import (
	"github.com/google/uuid"
)

type OutletOwner struct {
	OutletOwnerID uuid.UUID `json:"outlet_owner_id" gorm:"primaryKey;default:gen_random_uuid()"`
	OutletID      uuid.UUID `json:"outlet_id"`
	UserID        uuid.UUID `json:"user_id"`
	CreatedAt     int       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     int       `json:"updated_at" gorm:"autoUpdateTime"`

	Outlet *Outlet `json:"outlet" gorm:"->"`
	User   *User   `json:"user" gorm:"->"`
}

func (OutletOwner) TableName() string {
	return "tr_outlet_owner"
}
