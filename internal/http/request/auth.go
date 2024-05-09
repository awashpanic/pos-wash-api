package request

import (
	"github.com/ffajarpratama/pos-wash-api/pkg/types"
	"github.com/google/uuid"
)

type Register struct {
	User   UserRegister   `json:"user" validate:"required"`
	Outlet OutletRegister `json:"outlet" validate:"required"`
}

type UserRegister struct {
	Name        string            `json:"name" validate:"required"`
	Email       string            `json:"email" validate:"required,email"`
	PhoneNumber types.PhoneNumber `json:"phone_number" validate:"required"`
	Password    types.Password    `json:"password" validate:"required,min=8"`
}

type OutletRegister struct {
	Name    string     `json:"name" validate:"required"`
	Address string     `json:"address" validate:"required"`
	LogoID  *uuid.UUID `json:"logo_id"`
}

type Login struct {
	Identifier string         `json:"identifier" validate:"required"`
	Password   types.Password `json:"password" validate:"required"`
}
