package request

import (
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/types"
	"github.com/google/uuid"
)

type ReqInsertUser struct {
	Name        string              `json:"name" validate:"required"`
	Gender      constant.UserGender `json:"gender" validate:"required,oneof=male female"`
	PhoneNumber types.PhoneNumber   `json:"phone_number" validate:"required"`
	Address     string              `json:"address" validate:"required"`
	Role        constant.UserRole   `json:"-"`
	OutletID    uuid.UUID           `json:"-"`
}
