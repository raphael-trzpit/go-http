package pagination

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Paginator struct {
	config Config
}

func NewPaginator(config Config) Paginator {
	return Paginator{config: config}
}

func (p Paginator) FromRequest(r *http.Request) *Pagination {
	return p.FromValues(r.URL.Query())
}

func (p Paginator) FromValues(v url.Values) *Pagination {
	return &Pagination{
		page:    p.getPage(v),
		perPage: p.getPerPage(v),
	}
}

func (p Paginator) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pagination := p.FromRequest(r)

		h.ServeHTTP(w, r.WithContext(NewContext(r.Context(), pagination)))

		urlRequested := r.URL
		linkHeader := make([]string, 0)
		for _, link := range pagination.GetLinks() {
			urlRequested.Query().Set(p.config.KeyPage, string(link.page))
			urlRequested.Query().Set(p.config.KeyPerPage, string(link.perPage))
			linkHeader = append(linkHeader, "<"+urlRequested.String()+"> rel=\""+link.rel.String()+"\"")
		}
		w.Header().Set("Link", strings.Join(linkHeader, ", "))
	})
}

func (p Paginator) getPage(v url.Values) uint {
	page, err := strconv.Atoi(v.Get(p.config.keyPage()))
	if err != nil {
		return p.config.page()
	}

	return uint(page)
}

func (p Paginator) getPerPage(v url.Values) uint {
	perPage, err := strconv.Atoi(v.Get(p.config.keyPerPage()))
	if err != nil || perPage <= 0 {
		return p.config.perPage()
	}

	return uint(perPage)
}
