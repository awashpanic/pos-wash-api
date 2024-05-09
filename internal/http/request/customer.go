package request

import (
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/types"
	"github.com/google/uuid"
)

type ListCustomerQuery struct {
	BaseQuery
	OutletID uuid.UUID
}

type CreateCustomer struct {
	Name        string              `json:"name" validate:"required"`
	PhoneNumber types.PhoneNumber   `json:"phone_number" validate:"required"`
	Gender      constant.UserGender `json:"gender" validate:"required,oneof=male female"`
	Address     string              `json:"address" validate:"required"`
	AvatarID    *uuid.UUID          `json:"avatar_id"`
	OutletID    uuid.UUID           `json:"-"`
}

type UpdateCustomer struct {
	Name        string              `json:"name" validate:"required"`
	PhoneNumber types.PhoneNumber   `json:"phone_number" validate:"required"`
	Gender      constant.UserGender `json:"gender" validate:"required,oneof=male female"`
	Address     string              `json:"address" validate:"required"`
	AvatarID    *uuid.UUID          `json:"avatar_id"`
	OutletID    uuid.UUID           `json:"-"`
	CustomerID  uuid.UUID           `json:"-"`
}
