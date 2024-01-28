package request

import "github.com/google/uuid"

type ListOutletQuery struct {
	BaseQuery
}

type ReqInsertOutlet struct {
	Name     string     `json:"name" validate:"required"`
	Address  string     `json:"address" validate:"required"`
	LogoID   *uuid.UUID `json:"logo_id"`
	OutletID uuid.UUID  `json:"-"`
	UserID   uuid.UUID  `json:"-"`
}
