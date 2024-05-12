package usecase

import (
	"context"
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/internal/repository"
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/custom_error"
	"github.com/google/uuid"
)

// CreateCustomer implements IFaceUsecase.
func (u *Usecase) CreateCustomer(ctx context.Context, req *request.CreateCustomer) error {
	customer := &model.Customer{
		OutletID:    req.OutletID,
		AvatarID:    req.AvatarID,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber.Format(),
		Gender:      req.Gender,
		Address:     req.Address,
	}

	err := u.Repo.CreateCustomer(ctx, customer, u.DB)
	if err != nil {
		if repository.IsDuplicateErr(err) {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusConflict,
				Code:     constant.DefaultDuplicateDataError,
				Message:  "phone already used",
			})
		}

		return err
	}

	return nil
}

// FindAndCountCustomer implements IFaceUsecase.
func (u *Usecase) FindAndCountCustomer(ctx context.Context, params *request.ListCustomerQuery) ([]*model.Customer, int64, error) {
	return u.Repo.FindAndCountCustomer(ctx, params)
}

// FindOneCustomer implements IFaceUsecase.
func (u *Usecase) FindOneCustomer(ctx context.Context, customerID uuid.UUID) (*model.Customer, error) {
	return u.Repo.FindOneCustomer(ctx, "customer_id = ?", customerID)
}

// UpdateCustomer implements IFaceUsecase.
func (u *Usecase) UpdateCustomer(ctx context.Context, req *request.UpdateCustomer) error {
	data := map[string]interface{}{
		"avatar_id":    req.AvatarID,
		"name":         req.Name,
		"phone_number": req.PhoneNumber.Format(),
		"gender":       req.Gender,
		"address":      req.Address,
	}

	err := u.Repo.UpdateCustomer(ctx, u.DB, data, "customer_id = ?", req.CustomerID)
	if err != nil {
		if repository.IsDuplicateErr(err) {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusConflict,
				Code:     constant.DefaultDuplicateDataError,
				Message:  "phone already used",
			})
		}

		return err
	}

	return nil
}

// DeleteCustomer implements IFaceUsecase.
func (u *Usecase) DeleteCustomer(ctx context.Context, customerID uuid.UUID) error {
	return u.Repo.DeleteCustomer(ctx, u.DB, "customer_id = ?", customerID)
}
