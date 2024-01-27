package constant

type JwtKey int

const (
	UserIDKey JwtKey = iota
	RoleKey

	FILE_UPLOAD_MAX_SIZE = 1024 * 1024 * 10   // 10MB
	FILE_UPLOAD_MAX_AGE  = 365 * 24 * 60 * 60 // 1 year
)
