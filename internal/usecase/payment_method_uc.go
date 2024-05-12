package usecase

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
)

// FindAndCountPaymentMethod implements IFaceUsecase.
func (u *Usecase) FindAndCountPaymentMethod(ctx context.Context, params *request.ListPaymentMethodQuery) ([]*model.PaymentMethod, int64, error) {
	return u.Repo.FindAndCountPaymentMethod(ctx, params)
}
