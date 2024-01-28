package handler

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/google/uuid"
)

func (h *handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req request.ReqInsertUser
	err := h.v.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.Role = constant.Customer
	req.OutletID, _ = uuid.Parse(util.GetOutletIDFromCtx(ctx))

	err = h.uc.CreateUser(ctx, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, nil)
}

func (h *handler) FindAndCounCustomer(w http.ResponseWriter, r *http.Request) {
	params := new(request.ListUserOutletQuery)
	params.BaseQuery = request.NewBaseQuery(r)

	res, cnt, err := h.uc.FindAndCountUserOutlet(r.Context(), params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Paging(w, res, params.Page, params.Limit, cnt)
}
