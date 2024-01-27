package constant

type UserRole string
type UserGender string

const (
	SuperAdmin UserRole = "administrator"
	Owner      UserRole = "owner"
	Cashier    UserRole = "cashier"
	Customer   UserRole = "customer"

	Male   UserGender = "male"
	Female UserGender = "female"
)
