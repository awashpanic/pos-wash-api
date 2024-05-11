package response

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
	Yesterday float64 `json:"yesterday"`
	Today     float64 `json:"today"`
	Delta     float64 `json:"delta"`
}

type CustomerTrend struct {
	Yesterday float64 `json:"yesterday"`
	Today     float64 `json:"today"`
	Delta     float64 `json:"delta"`
}
