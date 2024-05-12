package usecase

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
)

// FindAndCountPerfume implements IFaceUsecase.
func (u *Usecase) FindAndCountPerfume(ctx context.Context, params *request.ListPerfumeQuery) ([]*model.Perfume, int64, error) {
	return u.Repo.FindAndCountPerfume(ctx, params)
}
