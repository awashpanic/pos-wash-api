package constant

type UserRole string
type UserGender string

const (
	SuperAdmin UserRole = "super-admin"
	Owner      UserRole = "owner"
	Cashier    UserRole = "cashier"
	Customer   UserRole = "customer"

	Male   UserGender = "male"
	Female UserGender = "female"
)
