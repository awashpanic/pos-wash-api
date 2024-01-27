package handler

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/google/uuid"
)

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	var req request.ReqRegister
	err := h.v.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	res, err := h.uc.Register(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var req request.ReqLogin
	err := h.v.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	res, err := h.uc.Login(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}

func (h *handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, _ := uuid.Parse(util.GetUserIDFromCtx(ctx))

	res, err := h.uc.FindOneUser(ctx, userID)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.OK(w, res)
}
