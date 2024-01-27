package hash

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	custom_error "github.com/ffajarpratama/pos-wash-api/pkg/error"
	"golang.org/x/crypto/bcrypt"
)

const (
	BCRYPT_COST = 11
)

func Compare(hashedPwd string, plainPwd []byte) (err error) {
	byteHash := []byte(hashedPwd)
	err = bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			Code:     constant.DefaultBadRequestError,
			HTTPCode: http.StatusBadRequest,
			Message:  "password salah",
		})

		return
	}

	return
}

func HashAndSalt(pwd []byte) (hashPwd string, err error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, BCRYPT_COST)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
