package util

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
)

func GetTokenFromHeader(r *http.Request) (token string, err error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		err = errors.New("token is empty")
		return "", err
	}

	lenToken := 2
	s := strings.Split(authHeader, " ")
	if len(s) != lenToken {
		err = errors.New("token is invalid")
		return "", err
	}

	token = s[1]
	return token, nil
}

func GetUserIDFromCtx(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if userID, ok := ctx.Value(constant.UserIDKey).(string); ok {
		return userID
	}

	return ""
}

func GetRoleFromCtx(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	role, ok := ctx.Value(constant.RoleKey).(string)
	if ok {
		return role
	}

	return ""
}
