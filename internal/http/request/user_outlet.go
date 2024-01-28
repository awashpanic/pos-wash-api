package request

import "github.com/ffajarpratama/pos-wash-api/pkg/constant"

type ListUserOutletQuery struct {
	BaseQuery
	Role constant.UserRole
}
