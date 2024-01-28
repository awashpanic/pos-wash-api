package repository

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"gorm.io/gorm"
)

// CreateOutletOwner implements IFaceRepository.
func (r *Repository) CreateOutletOwner(ctx context.Context, data *model.OutletOwner, db *gorm.DB) error {
	return r.BaseRepository.Create(db.WithContext(ctx), data)
}

// FindOneOutletOwner implements IFaceRepository.
func (r *Repository) FindOneOutletOwner(ctx context.Context, query ...interface{}) (*model.OutletOwner, error) {
	var res *model.OutletOwner

	if err := r.BaseRepository.FindOne(r.db.WithContext(ctx).Preload("Outlet").Preload("User").Where(query[0], query[1:]...), &res); err != nil {
		return nil, err
	}

	return res, nil
}
