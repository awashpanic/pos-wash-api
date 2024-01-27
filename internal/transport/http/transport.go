package transport

import (
	"encoding/json"
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/config"
	"github.com/ffajarpratama/pos-wash-api/internal/handler"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/internal/middleware"
	"github.com/ffajarpratama/pos-wash-api/internal/usecase"
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/recover"
	"github.com/ffajarpratama/pos-wash-api/pkg/redis"
	"github.com/ffajarpratama/pos-wash-api/pkg/types"
	custom_validator "github.com/ffajarpratama/pos-wash-api/pkg/validator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	go_validator "github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	"github.com/sirupsen/logrus"
)

func NewHTTPHandler(cnf *config.Config, uc usecase.IFaceUsecase, redis redis.IFaceRedis) http.Handler {
	r := chi.NewRouter()
	validator := go_validator.New()
	enTrans := en.New()
	uni := ut.New(enTrans, enTrans)
	trans, _ := uni.GetTranslator("en")

	logger := logrus.New()
	logger.SetFormatter(&types.CustomJSONFormatter{})

	en_trans.RegisterDefaultTranslations(validator, trans)
	v := custom_validator.New(validator, trans)
	mw := middleware.Middleware{
		Redis:     redis,
		JWTConfig: cnf.JWTConfig,
	}

	r.Use(middleware.Logger(logger))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(response.CONTENT_TYPE_HEADER, response.CONTENT_TYPE_JSON)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Errors: &response.ErrorResponse{
				Code:    constant.DefaultNotFoundError,
				Status:  http.StatusNotFound,
				Message: "please check url",
			},
		})
	})

	r.Use(recover.RecoverWrap)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(response.CONTENT_TYPE_HEADER, response.CONTENT_TYPE_JSON)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Success: true,
			Data: map[string]interface{}{
				"app-name": "pos-wash-api",
			},
		})
	})

	r.Mount("/api/pos", handler.NewRouter(uc, v, mw)) // api/v1/pos/*

	return r
}
