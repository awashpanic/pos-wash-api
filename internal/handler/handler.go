package handler

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/middleware"
	"github.com/ffajarpratama/pos-wash-api/internal/usecase"
	custom_validator "github.com/ffajarpratama/pos-wash-api/pkg/validator"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	uc usecase.IFaceUsecase
	v  custom_validator.Validator
}

func NewRouter(uc usecase.IFaceUsecase, v custom_validator.Validator, mw middleware.Middleware) http.Handler {
	r := chi.NewRouter()
	h := handler{
		uc: uc,
		v:  v,
	}

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
		r.Route("/profile", func(r chi.Router) {
			r.Use(mw.AuthenticateUser())
			r.Get("/", h.GetProfile)
		})
	})

	return r
}
