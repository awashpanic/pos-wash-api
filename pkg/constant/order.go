package constant

type OrderStatus string

const (
	OrderAccepted      OrderStatus = "accepted"
	OrderOnProcess     OrderStatus = "on-process"
	OrderWaitingPickup OrderStatus = "waiting-pickup"
	OrderComplete      OrderStatus = "complete"
)
