package repository

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
)

// FindAndCountPerfume implements IFaceRepository.
func (r *Repository) FindAndCountPerfume(ctx context.Context, params *request.ListPerfumeQuery) ([]*model.Perfume, int64, error) {
	var res = make([]*model.Perfume, 0)
	var cnt int64

	query := r.db.
		WithContext(ctx).
		Model(&model.Perfume{})

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

// FindOnePerfume implements IFaceRepository.
func (r *Repository) FindOnePerfume(ctx context.Context, query ...interface{}) (*model.Perfume, error) {
	var res *model.Perfume

	if err := r.BaseRepository.FindOne(r.db.WithContext(ctx).Where(query[0], query[1:]...), &res); err != nil {
		return nil, err
	}

	return res, nil
}
