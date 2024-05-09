package model

import (
	"time"

	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	ServiceID         uuid.UUID                         `json:"service_id" gorm:"primaryKey; default:gen_random_uuid()"`
	OutletID          uuid.UUID                         `json:"outlet_id"`
	ServiceCategoryID uuid.UUID                         `json:"service_category_id"`
	Name              string                            `json:"name"`
	Description       string                            `json:"description"`
	Price             float64                           `json:"price"`
	EstCompletion     int                               `json:"est_completion"`
	EstCompletionUnit constant.ServiceEstCompletionUnit `json:"est_completion_unit"`
	CreatedAt         time.Time                         `json:"created_at"`
	UpdatedAt         time.Time                         `json:"updated_at"`
	DeletedAt         gorm.DeletedAt                    `json:"-" gorm:"column:deleted_at"`

	Outlet          *Outlet          `json:"outlet" gorm:"foreignKey:OutletID; references:OutletID"`
	ServiceCategory *ServiceCategory `json:"service_category" gorm:"foreignKey:ServiceCategoryID; references:ServiceCategoryID"`
}

func (Service) TableName() string {
	return "tr_service"
}
