package middleware

import (
	"context"
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	custom_jwt "github.com/ffajarpratama/pos-wash-api/pkg/jwt"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateUser(secret string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token, err := util.GetTokenFromHeader(r)
			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			resJwt, err := jwt.ParseWithClaims(token, &custom_jwt.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			customClaims, ok := resJwt.Claims.(*custom_jwt.CustomClaims)
			if !ok && !resJwt.Valid {
				response.UnauthorizedError(w)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, constant.UserIDKey, customClaims.ID)
			ctx = context.WithValue(ctx, constant.RoleKey, customClaims.Role)
			ctx = context.WithValue(ctx, constant.OutletIDKey, customClaims.OutletID)
			h.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func ParseWithoutVerified(tokenString string) *custom_jwt.CustomClaims {
	res, _, err := new(jwt.Parser).ParseUnverified(tokenString, &custom_jwt.CustomClaims{})
	if err != nil {
		return nil
	}

	claims, ok := res.Claims.(*custom_jwt.CustomClaims)
	if ok && claims.ID != "" {
		return claims
	}

	return nil
}
