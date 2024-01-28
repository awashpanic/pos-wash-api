package usecase

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/google/uuid"
)

type IFaceUsecase interface {
	// auth
	Register(ctx context.Context, req *request.ReqRegister) (*model.User, error)
	Login(ctx context.Context, req *request.ReqLogin) (*model.User, error)

	// user
	CreateUser(ctx context.Context, req *request.ReqInsertUser) error
	FindOneUser(ctx context.Context, userID uuid.UUID) (*model.User, error)

	// outlet
	CreateOutlet(ctx context.Context, req *request.ReqInsertOutlet) error
	FindOneOutlet(ctx context.Context, outletID uuid.UUID) (*model.Outlet, error)

	// user outlet
	FindAndCountUserOutlet(ctx context.Context, params *request.ListUserOutletQuery) ([]*model.UserOutlet, int64, error)
}
