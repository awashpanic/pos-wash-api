package custom_validator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/custom_error"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	go_validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	validate *go_validator.Validate
	trans    ut.Translator
}

type ValidatorError struct {
	Code    int      `json:"code"`
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func (e ValidatorError) Error() string {
	return e.Message
}

func New() Validator {
	v := go_validator.New()
	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(v, trans)

	return Validator{
		validate: v,
		trans:    trans,
	}
}

func (v *Validator) ValidateStruct(r *http.Request, data interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	err = json.Unmarshal(body, data)
	if err != nil {
		fmt.Println("err:json.Unmarshal()", err)
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			Code:     constant.DefaultBadRequestError,
			HTTPCode: http.StatusUnprocessableEntity,
			Message:  "please check your body request",
		})

		return err
	}

	err = v.validate.Struct(data)
	if err == nil {
		return nil

	}

	var message string
	var details = make([]string, 0)
	for _, field := range err.(go_validator.ValidationErrors) {
		message = field.Translate(v.trans)
		details = append(details, message)
	}

	err = ValidatorError{
		Code:    constant.DefaultBadRequestError,
		Status:  http.StatusBadRequest,
		Message: message,
		Details: details,
	}

	return err
}
