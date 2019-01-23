package pagination

import (
	"context"
)

type Pagination struct {
	page        uint
	perPage     uint
	returnCount uint
	totalCount  uint
	hasNext     bool
}

func (p *Pagination) SetReturnCount(count uint) {
	p.returnCount = count
}
func (p *Pagination) SetTotalCount(count uint) {
	p.totalCount = count
}
func (p *Pagination) SetHasNext(hasNext bool) {
	p.hasNext = hasNext
}

type LinkRel int

const (
	First LinkRel = iota
	Prev
	Next
	Last
)

type Link struct {
	rel     LinkRel
	page    uint
	perPage uint
}

func (p *Pagination) GetLinks() []Link {
	if p.returnCount == 0 {
		return []Link{}
	}
	links := make([]Link, 0)
}

func (p *Pagination) addPrevLink([]Link) []Link {
	if p.returnCount > 0 && p.page > 1 {

	}
}

func (p *Pagination) Page() uint {
	return p.page
}

func (p *Pagination) PerPage() uint {
	return p.perPage
}

func (p *Pagination) Offset() uint {
	if p.page == 0 {
		return 0
	}
	return (p.page - 1) * p.perPage
}

func (p *Pagination) Limit() uint {
	return p.page * p.perPage
}

type contextKey int

var paginationKey contextKey

func NewContext(ctx context.Context, p *Pagination) context.Context {
	return context.WithValue(ctx, paginationKey, p)
}

func FromContext(ctx context.Context) (*Pagination, bool) {
	p, ok := ctx.Value(paginationKey).(*Pagination)
	return p, ok
}
