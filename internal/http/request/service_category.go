package request

import "github.com/google/uuid"

type ListServiceCategoryQuery struct {
	BaseQuery
	OutletID uuid.UUID
}
