package usecase

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
)

// FindAndCountServiceCategory implements IFaceUsecase.
func (u *Usecase) FindAndCountServiceCategory(ctx context.Context, params *request.ListServiceCategoryQuery) ([]*model.ServiceCategory, int64, error) {
	return u.Repo.FindAndCountServiceCategory(ctx, params)
}
