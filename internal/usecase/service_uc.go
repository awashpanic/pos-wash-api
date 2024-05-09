package usecase

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/google/uuid"
)

// CreateService implements IFaceUsecase.
func (u *Usecase) CreateService(ctx context.Context, req *request.CreateService) error {
	service := &model.Service{
		OutletID:          req.OutletID,
		ServiceCategoryID: req.ServiceCategoryID,
		Name:              req.Name,
		Description:       req.Description,
		Price:             req.Price,
		EstCompletion:     req.EstCompletion,
		EstCompletionUnit: req.EstCompletionUnit,
	}

	return u.Repo.CreateService(ctx, service, u.DB)
}

// FindAndCountService implements IFaceUsecase.
func (u *Usecase) FindAndCountService(ctx context.Context, params *request.ListServiceQuery) ([]*model.Service, int64, error) {
	return u.Repo.FindAndCountService(ctx, params)
}

// FindOneService implements IFaceUsecase.
func (u *Usecase) FindOneService(ctx context.Context, serviceID uuid.UUID) (*model.Service, error) {
	return u.Repo.FindOneService(ctx, "service_id = ?", serviceID)
}

// UpdateService implements IFaceUsecase.
func (u *Usecase) UpdateService(ctx context.Context, req *request.UpdateService) error {
	data := map[string]interface{}{
		"service_category_id": req.ServiceCategoryID,
		"name":                req.Name,
		"description":         req.Description,
		"price":               req.Price,
		"est_completion":      req.EstCompletion,
		"est_completion_unit": req.EstCompletionUnit,
	}

	return u.Repo.UpdateService(ctx, u.DB, data, "service_id = ?", req.ServiceID)
}

// DeleteService implements IFaceUsecase.
func (u *Usecase) DeleteService(ctx context.Context, serviceID uuid.UUID) error {
	return u.Repo.DeleteService(ctx, u.DB, "service_id = ?", serviceID)
}
