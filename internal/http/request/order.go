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
	PerfumeID  uuid.UUID `json:"perfume_id" validate:"required"`
	CustomerID uuid.UUID `json:"customer_id" validate:"required"`
	Services   []Service `json:"services" validate:"required,gt=0,dive"`
	Note       string    `json:"note"`
	UserID     uuid.UUID `json:"-"`
	OutletID   uuid.UUID `json:"-"`
}

type Service struct {
	ServiceID uuid.UUID `json:"service_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required,min=1"`
}

type UpdateOrderStatus struct {
	Status  constant.OrderStatus `json:"status" validate:"required,oneof=on-process waiting-pickup complete"`
	OrderID uuid.UUID            `json:"-"`
	UserID  uuid.UUID            `json:"-"`
}

type OrderPayment struct {
	PaymentMethodID   uuid.UUID  `json:"payment_method_id" validate:"required"`
	PaymentAmount     float64    `json:"payment_amount"`
	PaymentEvidenceID *uuid.UUID `json:"payment_evidence_id"`
	OrderID           uuid.UUID  `json:"-"`
}
