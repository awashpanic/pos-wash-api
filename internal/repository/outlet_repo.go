package repository

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateOutlet implements IFaceRepository.
func (r *Repository) CreateOutlet(ctx context.Context, data *model.Outlet, db *gorm.DB) error {
	return r.BaseRepository.Create(db.WithContext(ctx), data)
}

// FindAndCountOutlet implements IFaceRepository.
func (r *Repository) FindAndCountOutlet(ctx context.Context, params *request.ListOutletQuery) ([]*model.Outlet, int64, error) {
	var res = make([]*model.Outlet, 0)
	var cnt int64

	query := r.db.
		WithContext(ctx).
		Model(&model.Outlet{})

	if err := query.Count(&cnt).Error; err != nil {
		return nil, 0, err
	}

	if params.Sort != "" {
		query = query.Order(util.TransformSortClause("created_at", params.Sort))
	}

	if err := query.
		Limit(params.Limit).
		Offset(util.CalculateOffset(params.Page, params.Limit)).
		Find(&res).
		Error; err != nil {
		return nil, 0, err
	}

	return res, cnt, nil
}

// FindOneOutlet implements IFaceRepository.
func (r *Repository) FindOneOutlet(ctx context.Context, query ...interface{}) (*model.Outlet, error) {
	var res *model.Outlet

	if err := r.BaseRepository.FindOne(r.db.WithContext(ctx).Where(query[0], query[1:]...), &res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateOutlet implements IFaceRepository.
func (r *Repository) UpdateOutlet(ctx context.Context, outletID uuid.UUID, data map[string]interface{}, db *gorm.DB) error {
	return db.WithContext(ctx).Model(&model.Outlet{}).Where("outlet_id = ?", outletID).Updates(data).Error
}

// DeleteOutlet implements IFaceRepository.
func (r *Repository) DeleteOutlet(ctx context.Context, outletID uuid.UUID, db *gorm.DB) error {
	return db.WithContext(ctx).Where("outlet_id = ?", outletID).Delete(&model.Outlet{}).Error
}
