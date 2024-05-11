package repository

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateService implements IFaceRepository.
func (r *Repository) CreateService(ctx context.Context, data *model.Service, db *gorm.DB) error {
	return r.BaseRepository.Create(db.WithContext(ctx), data)
}

// FindAndCountService implements IFaceRepository.
func (r *Repository) FindAndCountService(ctx context.Context, params *request.ListServiceQuery) ([]*model.Service, int64, error) {
	var res = make([]*model.Service, 0)
	var cnt int64

	query := r.db.
		WithContext(ctx).
		Model(&model.Service{}).
		Preload("ServiceCategory").
		Preload("Outlet")

	if params.OutletID != uuid.Nil {
		query = query.Where("outlet_id = ?", params.OutletID)
	}

	if params.Keyword != "" {
		query = query.Where("name ILIKE ?", "%"+params.Keyword+"%")
	}

	if params.ServiceCategoryID != uuid.Nil {
		query = query.Where("service_category_id = ?", params.ServiceCategoryID)
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

// FindService implements IFaceRepository.
func (r *Repository) FindService(ctx context.Context, query ...interface{}) ([]*model.Service, error) {
	var res = make([]*model.Service, 0)

	if err := r.db.
		WithContext(ctx).
		Model(&model.Service{}).
		Where(query[0], query[1:]...).
		Find(&res).
		Error; err != nil {
		return nil, err
	}

	return res, nil
}

// FindOneService implements IFaceRepository.
func (r *Repository) FindOneService(ctx context.Context, query ...interface{}) (*model.Service, error) {
	var res *model.Service

	if err := r.BaseRepository.FindOne(
		r.db.WithContext(ctx).
			Where(query[0], query[1:]...).
			Preload("ServiceCategory").
			Preload("Outlet"),
		&res,
	); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateService implements IFaceRepository.
func (r *Repository) UpdateService(ctx context.Context, db *gorm.DB, data map[string]interface{}, query ...interface{}) error {
	return db.WithContext(ctx).Model(&model.Service{}).Where(query[0], query[1:]...).Updates(data).Error
}

// DeleteService implements IFaceRepository.
func (r *Repository) DeleteService(ctx context.Context, db *gorm.DB, query ...interface{}) error {
	return db.WithContext(ctx).Delete(&model.Service{}, query...).Error
}
