package request

import (
	"net/http"
	"strconv"
)

type BaseQuery struct {
	Page    int    `query:"page"`
	PerPage int    `query:"perPage"`
	Keyword string `query:"keyword"`
	Sort    string `query:"sort"`
}

const (
	DEFAULT_PAGE     = 1
	DEFAULT_PER_PAGE = 20
)

func NewBaseQuery(r *http.Request) BaseQuery {
	page := DEFAULT_PAGE
	limit := DEFAULT_PER_PAGE

	if r.URL.Query().Get("page") != "" {
		page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	}

	if r.URL.Query().Get("perPage") != "" {
		limit, _ = strconv.Atoi(r.URL.Query().Get("perPage"))
	}

	return BaseQuery{
		Page:    page,
		PerPage: limit,
		Keyword: r.URL.Query().Get("keyword"),
		Sort:    r.URL.Query().Get("sort"),
	}
}
