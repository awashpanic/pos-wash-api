package repository

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IFaceRepository interface {
	// user
	CreateUser(ctx context.Context, data *model.User, db *gorm.DB) error
	FindOneUser(ctx context.Context, query ...interface{}) (*model.User, error)

	// outlet
	CreateOutlet(ctx context.Context, data *model.Outlet, db *gorm.DB) error
	FindAndCountOutlet(ctx context.Context, params *request.ListOutletQuery) ([]*model.Outlet, int64, error)
	FindOneOutlet(ctx context.Context, query ...interface{}) (*model.Outlet, error)
	UpdateOutlet(ctx context.Context, outletID uuid.UUID, data map[string]interface{}, db *gorm.DB) error
	DeleteOutlet(ctx context.Context, outletID uuid.UUID, db *gorm.DB) error

	// service category
	FindAndCountServiceCategory(ctx context.Context, params *request.ListServiceCategoryQuery) ([]*model.ServiceCategory, int64, error)

	// service
	CreateService(ctx context.Context, data *model.Service, db *gorm.DB) error
	FindAndCountService(ctx context.Context, params *request.ListServiceQuery) ([]*model.Service, int64, error)
	FindOneService(ctx context.Context, query ...interface{}) (*model.Service, error)
	UpdateService(ctx context.Context, db *gorm.DB, data map[string]interface{}, query ...interface{}) error
	DeleteService(ctx context.Context, db *gorm.DB, query ...interface{}) error

	// customer
	CreateCustomer(ctx context.Context, data *model.Customer, db *gorm.DB) error
	FindAndCountCustomer(ctx context.Context, params *request.ListCustomerQuery) ([]*model.Customer, int64, error)
	FindOneCustomer(ctx context.Context, query ...interface{}) (*model.Customer, error)
	UpdateCustomer(ctx context.Context, db *gorm.DB, data map[string]interface{}, query ...interface{}) error
	DeleteCustomer(ctx context.Context, db *gorm.DB, query ...interface{}) error

	// perfume
	FindAndCountPerfume(ctx context.Context, params *request.ListPerfumeQuery) ([]*model.Perfume, int64, error)

	// payment method
	FindAndCountPaymentMethod(ctx context.Context, params *request.ListPaymentMethodQuery) ([]*model.PaymentMethod, int64, error)
	FindOnePaymentMethod(ctx context.Context, query ...interface{}) (*model.PaymentMethod, error)
}
