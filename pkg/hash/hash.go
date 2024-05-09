package hash

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/custom_error"
	"golang.org/x/crypto/bcrypt"
)

const (
	BCRYPT_COST = 11
)

func Compare(pwd []byte, str []byte) error {
	err := bcrypt.CompareHashAndPassword(pwd, str)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			Code:     constant.DefaultBadRequestError,
			HTTPCode: http.StatusBadRequest,
			Message:  "password salah",
		})

		return err
	}

	return nil
}

func HashAndSalt(pwd []byte) (hashPwd string, err error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, BCRYPT_COST)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
