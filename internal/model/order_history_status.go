package model

import (
	"time"

	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderHistoryStatus struct {
	OrderHistoryStatusID uuid.UUID            `json:"order_history_status_id" gorm:"primaryKey; default:gen_random_uuid()"`
	OrderID              uuid.UUID            `json:"order_id"`
	Status               constant.OrderStatus `json:"status"`
	CreatedBy            uuid.UUID            `json:"created_by"`
	CreatedAt            time.Time            `json:"created_at"`
	UpdatedAt            time.Time            `json:"updated_at"`
	DeletedAt            gorm.DeletedAt       `json:"-" gorm:"column:deleted_at"`

	CreatedByUser *User `json:"created_by_user" gorm:"foreignKey:CreatedBy; references:UserID"`
}

func (OrderHistoryStatus) TableName() string {
	return "tr_order_history_status"
}
