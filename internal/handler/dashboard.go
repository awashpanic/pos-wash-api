package handler

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/google/uuid"
)

func (h *handler) GetDashboardSummary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	outletID, _ := uuid.Parse(util.GetOutletIDFromCtx(ctx))
	res, err := h.uc.GetDashboardSummary(ctx, outletID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}

func (h *handler) GetOrderTrend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var params request.OrderTrendQuery
	params.OutletID, _ = uuid.Parse(util.GetOutletIDFromCtx(ctx))
	params.Start = r.URL.Query().Get("start")
	params.End = r.URL.Query().Get("end")
	params.Type = r.URL.Query().Get("type")

	res, err := h.uc.GetOrderTrend(ctx, &params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, res)
}
