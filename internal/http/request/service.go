package request

import (
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/google/uuid"
)

type ListServiceQuery struct {
	BaseQuery
	ServiceCategoryID uuid.UUID
	OutletID          uuid.UUID
}

type CreateService struct {
	Name              string                            `json:"name" validate:"required"`
	ServiceCategoryID uuid.UUID                         `json:"service_category_id" validate:"required"`
	Description       string                            `json:"description" validate:"required"`
	Price             float64                           `json:"price" validate:"required"`
	EstCompletion     int                               `json:"est_completion" validate:"required"`
	EstCompletionUnit constant.ServiceEstCompletionUnit `json:"est_completion_unit" validate:"required,oneof=minute hour day week month year"`
	OutletID          uuid.UUID                         `json:"-"`
}

type UpdateService struct {
	Name              string                            `json:"name" validate:"required"`
	ServiceCategoryID uuid.UUID                         `json:"service_category_id" validate:"required"`
	Description       string                            `json:"description" validate:"required"`
	Price             float64                           `json:"price" validate:"required"`
	EstCompletion     int                               `json:"est_completion" validate:"required"`
	EstCompletionUnit constant.ServiceEstCompletionUnit `json:"est_completion_unit" validate:"required,oneof=minute hour day week month year"`
	OutletID          uuid.UUID                         `json:"-"`
	ServiceID         uuid.UUID                         `json:"-"`
}
