package usecase

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/internal/repository"
	"github.com/google/uuid"
)

// FindOneUser implements IFaceUsecase.
func (u *Usecase) FindOneUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	res, err := u.Repo.FindOneUser(ctx, "user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	outletOwner, err := u.Repo.FindOneOutletOwner(ctx, "user_id = ?", res.UserID)
	if err != nil && !repository.IsRecordNotfound(err) {
		return nil, err
	}

	if outletOwner != nil {
		res.Outlet = outletOwner.Outlet
	}

	return res, nil
}
