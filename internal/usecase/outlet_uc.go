package usecase

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/google/uuid"
)

// CreateOutlet implements IFaceUsecase.
func (u *Usecase) CreateOutlet(ctx context.Context, req *request.ReqInsertOutlet) error {
	tx := u.DB.Begin()
	defer tx.Rollback()

	outlet := &model.Outlet{
		Name:    req.Name,
		Address: req.Address,
		LogoID:  req.LogoID,
	}

	err := u.Repo.CreateOutlet(ctx, outlet, tx)
	if err != nil {
		return err
	}

	outletOwner := &model.OutletOwner{
		OutletID: outlet.OutletID,
		UserID:   req.UserID,
	}

	err = u.Repo.CreateOutletOwner(ctx, outletOwner, tx)
	if err != nil {
		return err
	}

	return tx.Commit().Error
}

// FindAndCountOutlet implements IFaceUsecase.
func (u *Usecase) FindAndCountOutlet(ctx context.Context, params *request.ListOutletQuery) ([]*model.Outlet, int64, error) {
	return u.Repo.FindAndCountOutlet(ctx, params)
}

// FindOneOutlet implements IFaceUsecase.
func (u *Usecase) FindOneOutlet(ctx context.Context, outletID uuid.UUID) (*model.Outlet, error) {
	return u.Repo.FindOneOutlet(ctx, "outlet_id = ?", outletID)
}

// UpdateOutlet implements IFaceUsecase.
func (u *Usecase) UpdateOutlet(ctx context.Context, req *request.ReqInsertOutlet) error {
	panic("unimplemented")
}

// DeleteOutlet implements IFaceUsecase.
func (u *Usecase) DeleteOutlet(ctx context.Context, outletID uuid.UUID) error {
	return u.Repo.DeleteOutlet(ctx, outletID, u.DB)
}
