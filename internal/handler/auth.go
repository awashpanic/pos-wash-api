package handler

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
)

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	response.OK(w, nil)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	response.OK(w, nil)
}

func (h *handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	response.OK(w, nil)
}
