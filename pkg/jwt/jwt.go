package jwt

import (
	"time"

	"github.com/ffajarpratama/pos-wash-api/config"
	"github.com/golang-jwt/jwt/v5"
)

type CustomUserClaims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	UserSub  string `json:"user_sub"`
	AdminSub string `json:"admin_sub"`
	jwt.RegisteredClaims
}

type JWTRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

const (
	ACCESS_TTL  = time.Duration(1 * 24 * time.Hour)
	REFRESH_TTL = time.Duration(14 * 24 * time.Hour)
)

func GenerateToken(claims *CustomUserClaims, cnf *config.Config) (*JWTRes, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ACCESS_TTL)),
		Issuer:    "awash-panic.com",
	}

	accToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	access, err := accToken.SignedString([]byte(cnf.JWTConfig.User))
	if err != nil {
		return nil, err
	}

	refClaims := &RefreshTokenClaims{
		UserSub: claims.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(REFRESH_TTL)),
			Issuer:    "awash-panic.com",
		},
	}

	refToken := jwt.NewWithClaims(jwt.SigningMethodHS512, refClaims)
	refresh, err := refToken.SignedString([]byte(cnf.JWTConfig.Refresh))
	if err != nil {
		return nil, err
	}

	res := &JWTRes{
		AccessToken:  access,
		RefreshToken: refresh,
	}

	return res, nil
}

func GenerateTokenAdmin(claims *CustomUserClaims, secret string) (token string, err error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ACCESS_TTL)),
		Issuer:    "awash-panic.com",
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return jwtClaims.SignedString([]byte(secret))
}
