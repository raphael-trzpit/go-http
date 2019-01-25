package pagination

import (
	"context"
)

type Pagination struct {
	page       uint
	perPage    uint
	totalCount uint
	hasNext    bool
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

func (p *Pagination) SetTotalCount(count uint) {
	p.totalCount = count
}
func (p *Pagination) SetHasNext(hasNext bool) {
	p.hasNext = hasNext
}
func (p *Pagination) GetLinks() []Link {
	links := make([]Link, 0)
	links = p.addNextLinks(links)
	links = p.addPreviousLinks(links)
	links = p.addFirstLinks(links)
	links = p.addLastLink(links)

	return links
}

func (p *Pagination) addNextLinks(links []Link) []Link {
	if p.perPage == 0 {
		return links
	}

	if p.hasNext || p.Limit() < p.totalCount {
		return append(links, newLink(p, RelNext, p.page+1))
	}

	return links
}

func (p *Pagination) addPreviousLinks(links []Link) []Link {
	if p.page == 1 || p.perPage == 0 {
		return links
	}

	return append(links, newLink(p, RelPrev, p.page-1))
}

func (p *Pagination) addFirstLinks(links []Link) []Link {
	if p.perPage == 0 {
		return links
	}

	return append(links, newLink(p, RelFirst, 1))
}

func (p *Pagination) addLastLink(links []Link) []Link {
	if p.totalCount == 0 || p.perPage == 0 {
		return links
	}

	lastPage := p.totalCount / p.perPage
	if p.totalCount%p.perPage > 0 {
		lastPage += 1
	}

	return append(links, newLink(p, RelLast, lastPage))
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

type LinkRel int

const (
	RelFirst LinkRel = iota
	RelPrev
	RelNext
	RelLast
)

func (l LinkRel) String() string {
	switch l {
	case RelFirst:
		return "first"
	case RelLast:
		return "last"
	case RelNext:
		return "next"
	case RelPrev:
		return "prev"
	default:
		return "unknown"
	}
}

type Link struct {
	rel     LinkRel
	page    uint
	perPage uint
}

func newLink(p *Pagination, rel LinkRel, page uint) Link {
	return Link{
		rel:     rel,
		page:    page,
		perPage: p.PerPage(),
	}
}
