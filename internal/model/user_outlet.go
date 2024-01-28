package model

import (
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/google/uuid"
)

type UserOutlet struct {
	UserOutletID uuid.UUID         `json:"user_outlet_id" gorm:"primaryKey;default:gen_random_uuid()"`
	OutletID     uuid.UUID         `json:"outlet_id"`
	UserID       uuid.UUID         `json:"user_id"`
	Role         constant.UserRole `json:"role"`
	CreatedAt    int               `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    int               `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    int               `json:"-" gorm:"column:deleted_at"`

	Outlet *Outlet `json:"outlet" gorm:"->"`
	User   *User   `json:"user" gorm:"->"`
}

func (UserOutlet) TableName() string {
	return "tr_user_outlet"
}
