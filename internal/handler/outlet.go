package handler

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *handler) CreateOutlet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req request.ReqInsertOutlet
	err := h.v.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.UserID, _ = uuid.Parse(util.GetUserIDFromCtx(ctx))

	err = h.uc.CreateOutlet(ctx, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}

func (h *handler) FindAndCountOutlet(w http.ResponseWriter, r *http.Request) {
	params := new(request.ListOutletQuery)
	params.BaseQuery = request.NewBaseQuery(r)

	res, cnt, err := h.uc.FindAndCountOutlet(r.Context(), params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Paging(w, res, params.Page, params.Limit, cnt)
}

func (h *handler) FindOneOutlet(w http.ResponseWriter, r *http.Request) {
	outletID, _ := uuid.Parse(chi.URLParam(r, "outletID"))
	res, err := h.uc.FindOneOutlet(r.Context(), outletID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}

func (h *handler) UpdateOutlet(w http.ResponseWriter, r *http.Request) {
	response.OK(w, nil)
}

func (h *handler) DeleteOutlet(w http.ResponseWriter, r *http.Request) {
	outletID, _ := uuid.Parse(chi.URLParam(r, "outletID"))
	err := h.uc.DeleteOutlet(r.Context(), outletID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}
