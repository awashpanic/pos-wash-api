package usecase

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
)

// FindAndCountUserOutlet implements IFaceUsecase.
func (u *Usecase) FindAndCountUserOutlet(ctx context.Context, params *request.ListUserOutletQuery) ([]*model.UserOutlet, int64, error) {
	return u.Repo.FindAndCountUserOutlet(ctx, params)
}
