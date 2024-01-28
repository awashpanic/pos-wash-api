package repository

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mock/mock_repo.go -package=mock_repo -source=repo_interface.go
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

	// [!temp] outlet-owner
	CreateOutletOwner(ctx context.Context, data *model.OutletOwner, db *gorm.DB) error
	FindOneOutletOwner(ctx context.Context, query ...interface{}) (*model.OutletOwner, error)
}
