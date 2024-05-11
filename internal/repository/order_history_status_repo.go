package repository

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"gorm.io/gorm"
)

// CreateManyOrderHistoryStatus implements IFaceRepository.
func (r *Repository) CreateManyOrderHistoryStatus(ctx context.Context, data []*model.OrderHistoryStatus, db *gorm.DB) error {
	if len(data) == 0 {
		return nil
	}

	return r.BaseRepository.Create(db.WithContext(ctx), data)
}
