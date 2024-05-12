package response

import "github.com/ffajarpratama/pos-wash-api/internal/model"

type DashoardSummary struct {
	OrderCount    *OrderCount    `json:"order_count"`
	RevenueTrend  *RevenueTrend  `json:"revenue_trend"`
	CustomerTrend *CustomerTrend `json:"customer_trend"`
}

type OrderCount struct {
	Accepted  int64 `json:"accepted"`
	OnProcess int64 `json:"on_process"`
	Complete  int64 `json:"complete"`
}

type RevenueTrend struct {
	Initial float64 `json:"initial"`
	Final   float64 `json:"final"`
	Delta   float64 `json:"delta"`
}

type CustomerTrend struct {
	Initial float64 `json:"initial"`
	Final   float64 `json:"final"`
	Delta   float64 `json:"delta"`
}

type OrderTrend struct {
	OrderCount   *OrderCount         `json:"order_count"`
	RevenueTrend *RevenueTrend       `json:"revenue_trend"`
	Trend        []*model.OrderTrend `json:"trend"`
}
