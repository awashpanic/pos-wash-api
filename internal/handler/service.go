package handler

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *handler) CreateService(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req request.CreateService
	err := h.v.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.OutletID, _ = uuid.Parse(util.GetOutletIDFromCtx(ctx))

	err = h.uc.CreateService(ctx, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}

func (h *handler) FindAndCountService(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var params request.ListServiceQuery
	params.BaseQuery = *request.NewBaseQuery(r)
	params.OutletID, _ = uuid.Parse(util.GetOutletIDFromCtx(ctx))
	params.ServiceCategoryID, _ = uuid.Parse(r.URL.Query().Get("service_category_id"))

	res, cnt, err := h.uc.FindAndCountService(ctx, &params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Paging(w, res, params.Page, params.PerPage, cnt)
}

func (h *handler) FindOneService(w http.ResponseWriter, r *http.Request) {
	serviceID, _ := uuid.Parse(chi.URLParam(r, "serviceID"))
	res, err := h.uc.FindOneService(r.Context(), serviceID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}

func (h *handler) UpdateService(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req request.UpdateService
	err := h.v.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.ServiceID, _ = uuid.Parse(chi.URLParam(r, "serviceID"))
	req.OutletID, _ = uuid.Parse(util.GetOutletIDFromCtx(ctx))

	err = h.uc.UpdateService(ctx, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}

func (h *handler) DeleteService(w http.ResponseWriter, r *http.Request) {
	serviceID, _ := uuid.Parse(chi.URLParam(r, "serviceID"))
	err := h.uc.DeleteService(r.Context(), serviceID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}
