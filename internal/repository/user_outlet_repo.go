package repository

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"gorm.io/gorm"
)

// CreateUserOutlet implements IFaceRepository.
func (r *Repository) CreateUserOutlet(ctx context.Context, data *model.UserOutlet, db *gorm.DB) error {
	return r.BaseRepository.Create(db.WithContext(ctx), data)
}

// FindAndCountUserOutlet implements IFaceRepository.
func (r *Repository) FindAndCountUserOutlet(ctx context.Context, params *request.ListUserOutletQuery) ([]*model.UserOutlet, int64, error) {
	var res = make([]*model.UserOutlet, 0)
	var cnt int64

	query := r.db.
		WithContext(ctx).
		Model(&model.UserOutlet{}).
		Preload("User").
		Joins("JOIN tr_user ON tr_user.user_id = tr_user_outlet.user_id")

	if params.Keyword != "" {
		query = query.Where("tr_user.name ILIKE ?", "%"+params.Keyword+"%")
	}

	if params.Role != "" {
		query = query.Where("tr_user_outlet.role = ?", params.Role)
	}

	if err := query.Count(&cnt).Error; err != nil {
		return nil, 0, err
	}

	if params.Sort != "" {
		query = query.Order(util.TransformSortClause("tr_user_outlet.created_at", params.Sort))
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
