package request

import (
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/google/uuid"
)

type ListOrderQuery struct {
	BaseQuery
	OutletID uuid.UUID
	Status   constant.OrderStatus
	Paid     string
}

type CreateOrder struct {
	PaymentMethodID uuid.UUID `json:"payment_method_id" validate:"required"`
	PerfumeID       uuid.UUID `json:"perfume_id" validate:"required"`
	CustomerID      uuid.UUID `json:"customer_id" validate:"required"`
	Services        []Service `json:"services" validate:"required,gt=0,dive"`
	Note            string    `json:"note"`
	UserID          uuid.UUID `json:"-"`
	OutletID        uuid.UUID `json:"-"`
}

type Service struct {
	ServiceID uuid.UUID `json:"service_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required,min=1"`
}
