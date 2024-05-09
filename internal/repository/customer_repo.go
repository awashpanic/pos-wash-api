package repository

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateCustomer implements IFaceRepository.
func (r *Repository) CreateCustomer(ctx context.Context, data *model.Customer, db *gorm.DB) error {
	return r.BaseRepository.Create(db.WithContext(ctx), data)
}

// FindAndCountCustomer implements IFaceRepository.
func (r *Repository) FindAndCountCustomer(ctx context.Context, params *request.ListCustomerQuery) ([]*model.Customer, int64, error) {
	var res = make([]*model.Customer, 0)
	var cnt int64

	query := r.db.
		WithContext(ctx).
		Model(&model.Customer{}).
		Preload("Outlet").
		Preload("Avatar")

	if params.OutletID != uuid.Nil {
		query = query.Where("outlet_id = ?", params.OutletID)
	}

	if params.Keyword != "" {
		query = query.Where("name ILIKE ?", "%"+params.Keyword+"%")
	}

	if err := query.Count(&cnt).Error; err != nil {
		return nil, 0, err
	}

	if params.Sort != "" {
		query = query.Order(util.TransformSortClause("created_at", params.Sort))
	}

	if err := query.
		Limit(params.PerPage).
		Offset(util.CalculateOffset(params.Page, params.PerPage)).
		Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, cnt, nil
}

// FindOneCustomer implements IFaceRepository.
func (r *Repository) FindOneCustomer(ctx context.Context, query ...interface{}) (*model.Customer, error) {
	var res *model.Customer

	if err := r.BaseRepository.FindOne(
		r.db.
			WithContext(ctx).
			Where(query[0], query[1:]...).
			Preload("Outlet").
			Preload("Avatar"),
		&res,
	); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateCustomer implements IFaceRepository.
func (r *Repository) UpdateCustomer(ctx context.Context, db *gorm.DB, data map[string]interface{}, query ...interface{}) error {
	return db.WithContext(ctx).Model(&model.Customer{}).Where(query[0], query[1:]...).Updates(data).Error
}

// DeleteCustomer implements IFaceRepository.
func (r *Repository) DeleteCustomer(ctx context.Context, db *gorm.DB, query ...interface{}) error {
	return db.WithContext(ctx).Delete(&model.Customer{}, query...).Error
}
