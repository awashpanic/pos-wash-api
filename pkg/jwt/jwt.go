package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	ID       string `json:"id"`
	OutletID string `json:"outlet_id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

const (
	ACCESS_TTL  = time.Duration(1 * 24 * time.Hour)
	REFRESH_TTL = time.Duration(14 * 24 * time.Hour)
)

func GenerateToken(claims *CustomClaims, secret string) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ACCESS_TTL)),
		Issuer:    "awash-panic.com",
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(secret))
}
