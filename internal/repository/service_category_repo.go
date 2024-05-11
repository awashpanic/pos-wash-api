package repository

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
)

// FindAndCountServiceCategory implements IFaceRepository.
func (r *Repository) FindAndCountServiceCategory(ctx context.Context, params *request.ListServiceCategoryQuery) ([]*model.ServiceCategory, int64, error) {
	var res = make([]*model.ServiceCategory, 0)
	var cnt int64

	query := r.db.
		WithContext(ctx).
		Model(&model.ServiceCategory{})

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
