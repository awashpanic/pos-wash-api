package handler

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
)

func (h *handler) FindAndCountPerfume(w http.ResponseWriter, r *http.Request) {
	var params request.ListPerfumeQuery
	params.BaseQuery = request.NewBaseQuery(r)

	res, cnt, err := h.uc.FindAndCountPerfume(r.Context(), &params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Paging(w, res, params.Page, params.PerPage, cnt)
}
