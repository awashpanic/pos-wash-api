package model

import (
	"time"

	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	CustomerID  uuid.UUID           `json:"customer_id" gorm:"primaryKey;default:gen_random_uuid()"`
	OutletID    uuid.UUID           `json:"outlet_id"`
	AvatarID    *uuid.UUID          `json:"avatar_id"`
	Name        string              `json:"name"`
	PhoneNumber types.PhoneNumber   `json:"phone_number"`
	Gender      constant.UserGender `json:"gender"`
	Address     string              `json:"address"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt      `json:"-" gorm:"column:deleted_at"`

	Outlet *Outlet `json:"outlet" gorm:"foreignKey:OutletID; references:OutletID"`
	Avatar *Media  `json:"avatar" gorm:"foreignKey:AvatarID; references:MediaID"`
}

func (Customer) TableName() string {
	return "tr_customer"
}
