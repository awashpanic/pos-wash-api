package repository

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/custom_error"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository struct{}

func (r *BaseRepository) Create(db *gorm.DB, params interface{}) (err error) {
	err = db.Model(params).Create(params).Error
	if err != nil {
		fmt.Printf("[error] %v\n", err.Error())
		return err
	}

	return
}

func (r *BaseRepository) Update(db *gorm.DB, params interface{}) (err error) {
	rows := db.Omit(clause.Associations).Model(params).Updates(params)
	err = rows.Error
	if err != nil {
		fmt.Printf("[error] %v\n", err.Error())
		return err
	}

	if rows.RowsAffected == 0 {
		nameField := reflect.TypeOf(params).Elem().Name()
		msg := ""
		if nameField != "" {
			msg = fmt.Sprintf("%s Not found", nameField)
		}

		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Code:     constant.DefaultNotFoundError,
			Message:  msg,
		})

		return err
	}

	return
}

func (r *BaseRepository) FindOne(db *gorm.DB, result interface{}) (err error) {
	err = db.First(result).Error
	if err != nil {
		fmt.Printf("[error] %v\n", err.Error())

		if err == gorm.ErrRecordNotFound {
			nameField := reflect.TypeOf(result).Elem().Name()
			msg := ""
			if nameField != "" {
				msg = fmt.Sprintf("%s Not found", nameField)
			}

			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusNotFound,
				Code:     constant.DefaultNotFoundError,
				Message:  msg,
			})
		}

		return err
	}

	return
}

func (r *BaseRepository) Delete(db *gorm.DB, params interface{}) (err error) {
	err = db.Model(params).Delete(params).Error
	if err != nil {
		fmt.Printf("[error] %v\n", err.Error())

		if err == gorm.ErrRecordNotFound {
			nameField := reflect.TypeOf(params).Elem().Name()
			msg := ""
			if nameField != "" {
				msg = fmt.Sprintf("%s Not found", nameField)
			}

			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusNotFound,
				Code:     constant.DefaultNotFoundError,
				Message:  msg,
			})
		}

		return err
	}

	return
}

func IsDuplicateErr(err error) bool {
	return strings.Contains(err.Error(), "duplicate")
}

func IsRecordNotfound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
