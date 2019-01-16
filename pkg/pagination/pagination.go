package pagination

import (
	"context"
)

type Pagination struct {
	page    int
	perPage int
}

func (p Pagination) Page() int {
	return p.page
}

func (p Pagination) PerPage() int {
	return p.perPage
}

func (p Pagination) Offset() int {
	if p.page == 0 {
		return 0
	}
	return (p.page - 1) * p.perPage
}

func (p Pagination) Limit() int {
	return p.page * p.perPage
}

type contextKey int

var paginationKey contextKey

func NewContext(ctx context.Context, p Pagination) context.Context {
	return context.WithValue(ctx, paginationKey, p)
}

func FromContext(ctx context.Context) (Pagination, bool) {
	p, ok := ctx.Value(paginationKey).(Pagination)
	return p, ok
}
