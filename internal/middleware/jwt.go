package middleware

import (
	"context"
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/config"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	custom_jwt "github.com/ffajarpratama/pos-wash-api/pkg/jwt"
	"github.com/ffajarpratama/pos-wash-api/pkg/redis"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware struct {
	Redis redis.IFaceRedis
	config.JWTConfig
}

func (m Middleware) AuthenticateUser() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token, err := util.GetTokenFromHeader(r)
			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			resJwt, err := jwt.ParseWithClaims(token, &custom_jwt.CustomUserClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(m.JWTConfig.User), nil
			})

			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			customClaims, ok := resJwt.Claims.(*custom_jwt.CustomUserClaims)
			if !ok && !resJwt.Valid {
				response.UnauthorizedError(w)
				return
			}

			// _, err = m.Redis.Get(fmt.Sprintf("auth:%s", customClaims.SessionID))
			// if err != nil {
			// 	response.UnauthorizedError(w)
			// 	return
			// }

			ctx := r.Context()
			ctx = context.WithValue(ctx, constant.UserIDKey, customClaims.ID)
			ctx = context.WithValue(ctx, constant.RoleKey, customClaims.Role)
			h.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func (m Middleware) AuthenticateAdmin() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token, err := util.GetTokenFromHeader(r)
			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			resJwt, err := jwt.ParseWithClaims(token, &custom_jwt.CustomUserClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(m.JWTConfig.Admin), nil
			})

			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			customClaims, ok := resJwt.Claims.(*custom_jwt.CustomUserClaims)
			if !ok && !resJwt.Valid {
				response.UnauthorizedError(w)
				return
			}

			// _, err = m.Redis.Get(fmt.Sprintf("auth:%s", customClaims.SessionID))
			// if err != nil {
			// 	response.UnauthorizedError(w)
			// 	return
			// }

			ctx := r.Context()
			ctx = context.WithValue(ctx, constant.UserIDKey, customClaims.ID)
			ctx = context.WithValue(ctx, constant.RoleKey, customClaims.Role)
			h.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func (m Middleware) AccessLimiter(role constant.UserRole) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if role != constant.UserRole(util.GetRoleFromCtx(ctx)) {
				response.UnauthorizedError(w)
				return
			}

			h.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func (m Middleware) OptionalAuth(cnf *config.Config) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token, err := util.GetTokenFromHeader(r)
			if err != nil {
				ctx := r.Context()
				ctx = context.WithValue(ctx, constant.UserIDKey, custom_jwt.CustomUserClaims{ID: ""})
				h.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			resJwt, err := jwt.ParseWithClaims(token, &custom_jwt.CustomUserClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(cnf.JWTConfig.Admin), nil
			})

			if err != nil {
				resJwt, err = jwt.ParseWithClaims(token, &custom_jwt.CustomUserClaims{}, func(t *jwt.Token) (interface{}, error) {
					return []byte(cnf.JWTConfig.User), nil
				})

				if err != nil {
					response.UnauthorizedError(w)
					return
				}
			}

			customClaims, ok := resJwt.Claims.(*custom_jwt.CustomUserClaims)
			if !ok && !resJwt.Valid {
				response.UnauthorizedError(w)
				return
			}

			// _, err = m.Redis.Get(fmt.Sprintf("auth:%s", customClaims.SessionID))
			// if err != nil {
			// 	response.UnauthorizedError(w)
			// 	return
			// }

			ctx := r.Context()
			ctx = context.WithValue(ctx, constant.UserIDKey, customClaims.ID)
			ctx = context.WithValue(ctx, constant.RoleKey, customClaims.Role)
			h.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
