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

	r.Route("/auth", func(auth chi.Router) {
		auth.Post("/register", h.Register)
		auth.Post("/login", h.Login)
		auth.Route("/profile", func(profile chi.Router) {
			profile.Use(mw.AuthenticateUser())
			profile.Get("/", h.GetProfile)
		})
	})

	r.Group(func(pvt chi.Router) {
		pvt.Use(mw.AuthenticateUser())

		pvt.Route("/outlet", func(outlet chi.Router) {
			outlet.Post("/", h.CreateOutlet)
			outlet.Get("/", h.FindAndCountOutlet)
			outlet.Get("/{outletID}", h.FindOneOutlet)
			outlet.Put("/{outletID}", h.UpdateOutlet)
			outlet.Delete("/{outletID}", h.DeleteOutlet)
		})
	})

	return r
}
