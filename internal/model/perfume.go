package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Perfume struct {
	PerfumeID uuid.UUID      `json:"perfume_id" gorm:"primaryKey; default:gen_random_uuid()"`
	Name      string         `json:"name"`
	Price     float64        `json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

func (Perfume) TableName() string {
	return "tm_perfume"
}
