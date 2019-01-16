package pagination

import (
	"net/http"
	"net/url"
	"strconv"
)

type Paginator struct {
	config Config
}

func NewPaginator(config Config) Paginator {
	return Paginator{config: config}
}

func (p Paginator) FromRequest(r *http.Request) Pagination {
	return p.FromValues(r.URL.Query())
}

func (p Paginator) FromValues(v url.Values) Pagination {
	return Pagination{
		page: p.getPage(v),
		perPage: p.getPerPage(v),
	}
}

func (p Paginator) ParsingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newContext := NewContext(r.Context(), p.FromRequest(r))

		h.ServeHTTP(w, r.WithContext(newContext))
	})
}

func (p Paginator) getPage(v url.Values) int {
	page, err := strconv.Atoi(v.Get(p.config.keyPage()))
	if err != nil {
		return p.config.page()
	}

	return page
}

func (p Paginator) getPerPage(v url.Values) int {
	perPage, err := strconv.Atoi(v.Get(p.config.keyPerPage()))
	if err != nil {
		return p.config.perPage()
	}

	return perPage
}


