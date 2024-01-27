package usecase

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
)

type IFaceUsecase interface {
	// auth
	Register(ctx context.Context, req *request.ReqRegister) error
	Login(ctx context.Context, req *request.ReqLogin) error
}
