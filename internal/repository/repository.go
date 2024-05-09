package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	BaseRepository
	db *gorm.DB
}

func New(db *gorm.DB) IFaceRepository {
	return &Repository{
		db: db,
	}
}
