package repository

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
)

// FindAndCountPaymentMethod implements IFaceRepository.
func (r *Repository) FindAndCountPaymentMethod(ctx context.Context, params *request.ListPaymentMethodQuery) ([]*model.PaymentMethod, int64, error) {
	var res = make([]*model.PaymentMethod, 0)
	var cnt int64

	query := r.db.
		WithContext(ctx).
		Model(&model.PaymentMethod{})

	if params.Keyword != "" {
		query = query.Where("(name ILIKE ? OR label ILIKE ?)", "%"+params.Keyword+"%", "%"+params.Keyword+"%")
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

// FindOnePaymentMethod implements IFaceRepository.
func (r *Repository) FindOnePaymentMethod(ctx context.Context, query ...interface{}) (*model.PaymentMethod, error) {
	var res *model.PaymentMethod

	if err := r.BaseRepository.FindOne(r.db.WithContext(ctx).Where(query[0], query[1:]...), &res); err != nil {
		return nil, err
	}

	return res, nil
}
