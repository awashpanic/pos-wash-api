package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderDetail struct {
	OrderDetailID uuid.UUID      `json:"order_detail_id" gorm:"primaryKey; default:gen_random_uuid()"`
	OrderID       uuid.UUID      `json:"order_id"`
	OutletID      uuid.UUID      `json:"outlet_id"`
	ServiceID     uuid.UUID      `json:"service_id"`
	Quantity      int            `json:"quantity"`
	TotalAmount   float64        `json:"total_amount"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`

	Order   *Order   `json:"order" gorm:"foreignKey:OrderID; references:OrderID"`
	Outlet  *Outlet  `json:"outlet" gorm:"foreignKey:OutletID; references:OutletID"`
	Service *Service `json:"service" gorm:"foreignKey:ServiceID; references:ServiceID"`
}

func (OrderDetail) TableName() string {
	return "tr_order_detail"
}
