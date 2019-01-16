package pagination

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleFromContext() {
	// We use the default key config
	paginator := NewPaginator(Config{PerPage: 10})

	// Out HTTP handler that will retrieve the pagination
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pagination, _ := FromContext(r.Context())
		fmt.Println("page", pagination.Page())
		fmt.Println("perPage", pagination.PerPage())
		fmt.Println("offset", pagination.Offset())
		fmt.Println("limit", pagination.Limit())
	})

	// Use the paginator middleware
	paginator.ParsingMiddleware(handler).ServeHTTP(
		httptest.NewRecorder(),
		httptest.NewRequest(
			http.MethodGet,
			"http://host.com/page#1?custom=1&page=4&per_page=20",
			strings.NewReader(""),
		),
	)

	// Output:
	// page 4
	// perPage 20
	// offset 60
	// limit 80
}

func TestPaginator_ParsingMiddleware(t *testing.T) {
	tests := []struct {
		name               string
		config             Config
		target             string
		expectedPagination Pagination
	}{
		{
			name:               "Default config without params",
			config:             Config{},
			target:             "http://host.com/page",
			expectedPagination: Pagination{page: 1, perPage: 0},
		},
		{
			name:               "Default config with params",
			config:             Config{},
			target:             "http://host.com/page?page=2&per_page=15",
			expectedPagination: Pagination{page: 2, perPage: 15},
		},
		{
			name:               "Custom config without params",
			config:             Config{KeyPage: "customPage", KeyPerPage: "customPerPage", Page: 2, PerPage: 20},
			target:             "http://host.com/page?page=3&per_page=15",
			expectedPagination: Pagination{page: 2, perPage: 20},
		},
		{
			name:               "Custom config with params",
			config:             Config{KeyPage: "customPage", KeyPerPage: "customPerPage", Page: 2, PerPage: 20},
			target:             "http://host.com/page?page=2&per_page=15&customPage=3&customPerPage=30",
			expectedPagination: Pagination{page: 3, perPage: 30},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			paginator := NewPaginator(test.config)

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				pagination, _ := FromContext(r.Context())
				assert.Equal(t, test.expectedPagination, pagination)
			})

			// Use the paginator middleware
			paginator.ParsingMiddleware(handler).ServeHTTP(
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodGet, test.target, strings.NewReader("")),
			)
		})
	}
}
