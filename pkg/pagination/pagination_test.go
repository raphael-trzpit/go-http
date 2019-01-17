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
