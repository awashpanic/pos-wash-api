package router

import (
	"encoding/json"
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/config"
	"github.com/ffajarpratama/pos-wash-api/internal/handler"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/internal/middleware"
	"github.com/ffajarpratama/pos-wash-api/internal/usecase"
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/custom_validator"
	"github.com/go-chi/chi/v5"

	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func NewHTTPHandler(cnf *config.Config, uc usecase.IFaceUsecase) http.Handler {
	r := chi.NewRouter()
	v := custom_validator.New()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Error: &response.ErrorResponse{
				Code:    constant.DefaultNotFoundError,
				Status:  http.StatusNotFound,
				Message: "please check url",
			},
		})
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Error: &response.ErrorResponse{
				Code:    constant.DefaultMethodNotAllowed,
				Status:  http.StatusMethodNotAllowed,
				Message: "method not allowed",
			},
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Success: true,
			Data: map[string]interface{}{
				"app-name": "pos-wash-api",
			},
		})
	})

	r.Route("/api", func(api chi.Router) {
		api.Mount("/v1/pos", handler.NewV1Handler(cnf, uc, v))
	})

	return r
}
