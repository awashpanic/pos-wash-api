package handler

import (
	"net/http"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
)

func (h *handler) FindAndCountServiceCategory(w http.ResponseWriter, r *http.Request) {
	params := request.NewBaseQuery(r)
	res, cnt, err := h.uc.FindAndCountServiceCategory(r.Context(), params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Paging(w, res, params.Page, params.PerPage, cnt)
}
