package request

import (
	"github.com/google/uuid"
)

type OrderTrendQuery struct {
	OutletID uuid.UUID
	Start    string
	End      string
	Type     string
}
