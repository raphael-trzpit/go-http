package pagination

import (
	"context"
)

type Pagination struct {
	page    uint
	perPage uint
}

func (p Pagination) Page() uint {
	return p.page
}

func (p Pagination) PerPage() uint {
	return p.perPage
}

func (p Pagination) Offset() uint {
	if p.page == 0 {
		return 0
	}
	return (p.page - 1) * p.perPage
}

func (p Pagination) Limit() uint {
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
