package model

import (
	"time"

	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	OrderID          uuid.UUID            `json:"order_id" gorm:"primaryKey; default:gen_random_uuid()"`
	OutletID         uuid.UUID            `json:"outlet_id"`
	PaymentMethodID  *uuid.UUID           `json:"payment_method_id"`
	PerfumeID        uuid.UUID            `json:"perfume_id"`
	CustomerID       uuid.UUID            `json:"customer_id"`
	UserID           uuid.UUID            `json:"user_id"`
	InvoiceNumber    string               `json:"invoice_number"`
	SubtotalAmount   float64              `json:"subtotal_amount"`
	SubtotalFee      float64              `json:"subtotal_fee"`
	SubtotalDiscount float64              `json:"subtotal_discount"`
	TotalAmount      float64              `json:"total_amount"`
	Status           constant.OrderStatus `json:"status"`
	Note             string               `json:"note"`
	EstCompletionAt  time.Time            `json:"est_completion_at"`
	PaidAt           *time.Time           `json:"paid_at"`
	PaymentAmount    float64              `json:"payment_amount"`
	Change           float64              `json:"change"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        time.Time            `json:"updated_at"`
	DeletedAt        gorm.DeletedAt       `json:"-" gorm:"column:deleted_at"`

	Outlet             *Outlet               `json:"outlet" gorm:"foreignKey:OutletID; references:OutletID"`
	PaymentMethod      *PaymentMethod        `json:"payment_method" gorm:"foreignKey:PaymentMethodID; references:PaymentMethodID"`
	Perfume            *Perfume              `json:"perfume" gorm:"foreignKey:PerfumeID; references:PerfumeID"`
	Customer           *Customer             `json:"customer" gorm:"foreignKey:CustomerID; references:CustomerID"`
	User               *User                 `json:"user" gorm:"foreignKey:UserID; references:UserID"`
	OrderDetail        []*OrderDetail        `json:"order_detail" gorm:"foreignKey:OrderID; references:OrderID"`
	OrderHistoryStatus []*OrderHistoryStatus `json:"order_history_status" gorm:"foreignKey:OrderID; references:OrderID"`
}

func (Order) TableName() string {
	return "tr_order"
}

type OrderTrend struct {
	Accepted  int64   `gorm:"column:accepted"`
	OnProcess int64   `gorm:"column:on_process"`
	Complete  int64   `gorm:"column:complete"`
	Rev1      float64 `gorm:"column:rev_1"`
	Rev2      float64 `gorm:"column:rev_2"`
}
