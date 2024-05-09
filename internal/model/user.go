package model

import (
	"time"

	"github.com/ffajarpratama/pos-wash-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID      uuid.UUID         `json:"user_id" gorm:"primaryKey; default:gen_random_uuid()"`
	AvatarID    *uuid.UUID        `json:"avatar_id"`
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	PhoneNumber types.PhoneNumber `json:"phone_number"`
	Password    types.Password    `json:"-"`
	CreatedAt   time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time         `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt    `json:"-" gorm:"column:deleted_at"`

	Avatar *Media  `json:"avatar" gorm:"foreignKey:AvatarID; references:MediaID"`
	Outlet *Outlet `json:"outlet" gorm:"foreignKey:UserID; references:UserID"`

	// json fields
	AccessToken string `json:"access_token,omitempty" gorm:"-"`
}

func (User) TableName() string {
	return "tr_user"
}
