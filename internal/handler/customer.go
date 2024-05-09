package handler

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req request.CreateCustomer
	err := h.v.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.OutletID, _ = uuid.Parse(util.GetOutletIDFromCtx(ctx))

	err = h.uc.CreateCustomer(ctx, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}

func (h *handler) FindAndCountCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var params request.ListCustomerQuery
	params.BaseQuery = *request.NewBaseQuery(r)
	params.OutletID, _ = uuid.Parse(util.GetOutletIDFromCtx(ctx))

	res, cnt, err := h.uc.FindAndCountCustomer(ctx, &params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Paging(w, res, params.Page, params.PerPage, cnt)
}

func (h *handler) FindOneCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := uuid.Parse(chi.URLParam(r, "customerID"))
	res, err := h.uc.FindOneCustomer(r.Context(), customerID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}

func (h *handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req request.UpdateCustomer
	err := h.v.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.CustomerID, _ = uuid.Parse(chi.URLParam(r, "customerID"))
	req.OutletID, _ = uuid.Parse(util.GetOutletIDFromCtx(ctx))

	err = h.uc.UpdateCustomer(ctx, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}

func (h *handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := uuid.Parse(chi.URLParam(r, "customerID"))
	err := h.uc.DeleteCustomer(r.Context(), customerID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}
