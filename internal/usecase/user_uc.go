package usecase

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/internal/repository"
	"github.com/ffajarpratama/pos-wash-api/pkg/types"
	"github.com/google/uuid"
)

// CreateUser implements IFaceUsecase.
func (u *Usecase) CreateUser(ctx context.Context, req *request.ReqInsertUser) error {
	tx := u.DB.Begin()
	defer tx.Rollback()

	user := &model.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber.Format(),
		Password:    types.Password(req.PhoneNumber.Format()).Hash(), // customer's pass defaults to their phone_number
		Gender:      req.Gender,
		Address:     req.Address,
	}

	err := u.Repo.CreateUser(ctx, user, tx)
	if err != nil {
		return err
	}

	userOutlet := &model.UserOutlet{
		OutletID: req.OutletID,
		UserID:   user.UserID,
		Role:     req.Role,
	}

	err = u.Repo.CreateUserOutlet(ctx, userOutlet, tx)
	if err != nil {
		return err
	}

	return tx.Commit().Error
}

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
