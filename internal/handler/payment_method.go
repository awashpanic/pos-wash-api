package handler

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
)

func (h *handler) FindAndCountPaymentMethod(w http.ResponseWriter, r *http.Request) {
	var params request.ListPaymentMethodQuery
	params.BaseQuery = request.NewBaseQuery(r)

	res, cnt, err := h.uc.FindAndCountPaymentMethod(r.Context(), &params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Paging(w, res, params.Page, params.PerPage, cnt)
}
