package model

import (
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	UserID      uuid.UUID             `json:"user_id" gorm:"primaryKey;default:gen_random_uuid()"`
	OutletID    *uuid.UUID            `json:"outlet_id"`
	AvatarID    *uuid.UUID            `json:"avatar_id"`
	Name        string                `json:"name"`
	Email       string                `json:"email"`
	PhoneNumber types.PhoneNumber     `json:"phone_number"`
	Password    types.Password        `json:"-"`
	Role        constant.UserRole     `json:"role"`
	CreatedAt   int                   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int                   `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   soft_delete.DeletedAt `json:"-" gorm:"column:deleted_at"`

	Outlet *Outlet `json:"outlet" gorm:"->"`
	Avatar *Media  `json:"avatar" gorm:"foreignKey:AvatarID;references:MediaID"`

	// json fields
	AccessToken  string `json:"access_token,omitempty" gorm:"-"`
	RefreshToken string `json:"refresh_token,omitempty" gorm:"-"`
}

func (User) TableName() string {
	return "tr_user"
}
