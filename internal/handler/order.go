package handler

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req request.CreateOrder
	err := h.v.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.OutletID, _ = uuid.Parse(util.GetOutletIDFromCtx(ctx))
	req.UserID, _ = uuid.Parse(util.GetUserIDFromCtx(ctx))

	res, err := h.uc.CreateOrder(ctx, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}

func (h *handler) FindAndCountOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var params request.ListOrderQuery
	params.BaseQuery = request.NewBaseQuery(r)
	params.Status = constant.OrderStatus(r.URL.Query().Get("status"))
	params.Paid = r.URL.Query().Get("paid")
	params.OutletID, _ = uuid.Parse(util.GetOutletIDFromCtx(ctx))

	res, cnt, err := h.uc.FindAndCountOrder(ctx, &params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Paging(w, res, params.Page, params.PerPage, cnt)
}

func (h *handler) FindOneOrder(w http.ResponseWriter, r *http.Request) {
	orderID, _ := uuid.Parse(chi.URLParam(r, "orderID"))
	res, err := h.uc.FindOneOrder(r.Context(), orderID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}

func (h *handler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req request.UpdateOrderStatus
	err := h.v.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.OrderID, _ = uuid.Parse(chi.URLParam(r, "orderID"))
	req.UserID, _ = uuid.Parse(util.GetUserIDFromCtx(ctx))

	err = h.uc.UpdateOrderStatus(ctx, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}

func (h *handler) OrderPayment(w http.ResponseWriter, r *http.Request) {
	var req request.OrderPayment
	err := h.v.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.OrderID, _ = uuid.Parse(chi.URLParam(r, "orderID"))

	err = h.uc.OrderPayment(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}
