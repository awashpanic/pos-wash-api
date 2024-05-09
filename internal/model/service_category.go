package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServiceCategory struct {
	ServiceCategoryID uuid.UUID      `json:"service_category_id" gorm:"primaryKey; default:gen_random_uuid()"`
	Name              string         `json:"name"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

func (ServiceCategory) TableName() string {
	return "tm_service_category"
}
