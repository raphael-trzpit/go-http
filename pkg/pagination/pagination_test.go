package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagination(t *testing.T) {
	tests := []struct {
		name           string
		pagination     Pagination
		expectedOffset uint
		expectedLimit  uint
	}{
		{
			name: "zero value", pagination: Pagination{},
			expectedOffset: 0, expectedLimit: 0,
		},
		{
			name: "page 0", pagination: Pagination{page: 0, perPage: 10},
			expectedOffset: 0, expectedLimit: 0,
		},
		{
			name: "page 1", pagination: Pagination{page: 1, perPage: 10},
			expectedOffset: 0, expectedLimit: 10,
		},
		{
			name: "page 2", pagination: Pagination{page: 2, perPage: 10},
			expectedOffset: 10, expectedLimit: 20,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedOffset, test.pagination.Offset())
			assert.Equal(t, test.expectedLimit, test.pagination.Limit())
		})
	}
}

func TestPagination_GetLinks(t *testing.T) {
	tests := []struct {
		name       string
		pagination Pagination
		links      []Link
	}{
		{
			name: "exact count", pagination: Pagination{page: 3, perPage: 10, totalCount: 50},
			links: []Link{
				Link{rel: RelFirst, page: 1, perPage: 10},
				Link{rel: RelPrev, page: 2, perPage: 10},
				Link{rel: RelNext, page: 4, perPage: 10},
				Link{rel: RelLast, page: 5, perPage: 10},
			},
		},
		{
			name: "not exact count", pagination: Pagination{page: 3, perPage: 10, totalCount: 48},
			links: []Link{
				Link{rel: RelFirst, page: 1, perPage: 10},
				Link{rel: RelPrev, page: 2, perPage: 10},
				Link{rel: RelNext, page: 4, perPage: 10},
				Link{rel: RelLast, page: 5, perPage: 10},
			},
		},
		{
			name: "without totalCount", pagination: Pagination{page: 3, perPage: 10},
			links: []Link{
				Link{rel: RelFirst, page: 1, perPage: 10},
				Link{rel: RelPrev, page: 2, perPage: 10},
			},
		},
		{
			name: "without totalCount but with hasNext", pagination: Pagination{page: 3, perPage: 10, hasNext: true},
			links: []Link{
				Link{rel: RelFirst, page: 1, perPage: 10},
				Link{rel: RelPrev, page: 2, perPage: 10},
				Link{rel: RelNext, page: 4, perPage: 10},
			},
		},
		{
			name: "first and prev are the same", pagination: Pagination{page: 2, perPage: 10, totalCount: 40},
			links: []Link{
				Link{rel: RelFirst, page: 1, perPage: 10},
				Link{rel: RelPrev, page: 1, perPage: 10},
				Link{rel: RelNext, page: 3, perPage: 10},
				Link{rel: RelLast, page: 4, perPage: 10},
			},
		},
		{
			name: "next and last are the same", pagination: Pagination{page: 3, perPage: 10, totalCount: 40},
			links: []Link{
				Link{rel: RelFirst, page: 1, perPage: 10},
				Link{rel: RelPrev, page: 2, perPage: 10},
				Link{rel: RelNext, page: 4, perPage: 10},
				Link{rel: RelLast, page: 4, perPage: 10},
			},
		},
		{
			name: "on first page", pagination: Pagination{page: 1, perPage: 10, totalCount: 40},
			links: []Link{
				Link{rel: RelFirst, page: 1, perPage: 10},
				Link{rel: RelNext, page: 2, perPage: 10},
				Link{rel: RelLast, page: 4, perPage: 10},
			},
		},
		{
			name: "on last page", pagination: Pagination{page: 4, perPage: 10, totalCount: 40},
			links: []Link{
				Link{rel: RelFirst, page: 1, perPage: 10},
				Link{rel: RelPrev, page: 3, perPage: 10},
				Link{rel: RelLast, page: 4, perPage: 10},
			},
		},
		{
			name: "without perPage", pagination: Pagination{page: 4, totalCount: 40},
			links: []Link{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.ElementsMatch(t, test.links, test.pagination.GetLinks())
		})
	}
}
