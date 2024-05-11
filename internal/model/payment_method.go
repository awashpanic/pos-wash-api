package model

import (
	"time"

	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	PaymentMethodID uuid.UUID               `json:"payment_method_id" gorm:"primaryKey; default:gen_random_uuid()"`
	MediaID         *uuid.UUID              `json:"media_id"`
	Name            string                  `json:"name"`
	Label           string                  `json:"label"`
	Channel         constant.PaymentChannel `json:"channel"`
	Fee             float64                 `json:"fee"`
	FeePercentage   float64                 `json:"fee_percentage"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
	DeletedAt       gorm.DeletedAt          `json:"-" gorm:"column:deleted_at"`

	Media *Media `json:"media" gorm:"foreignKey:MediaID; references:MediaID"`
}

func (PaymentMethod) TableName() string {
	return "tm_payment_method"
}
