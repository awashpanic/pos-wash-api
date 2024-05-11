package repository

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"gorm.io/gorm"
)

// CreateManyOrderDetail implements IFaceRepository.
func (r *Repository) CreateManyOrderDetail(ctx context.Context, data []*model.OrderDetail, db *gorm.DB) error {
	if len(data) == 0 {
		return nil
	}

	return r.BaseRepository.Create(db.WithContext(ctx), data)
}
