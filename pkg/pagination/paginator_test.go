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
		pagination.SetTotalCount(119)
		fmt.Println("page", pagination.Page())
		fmt.Println("perPage", pagination.PerPage())
		fmt.Println("offset", pagination.Offset())
		fmt.Println("limit", pagination.Limit())
	})

	// Use the paginator middleware
	recorder := httptest.NewRecorder()
	paginator.Middleware(handler).ServeHTTP(
		recorder,
		httptest.NewRequest(
			http.MethodGet,
			"http://host.com/page#anchor?custom=1&page=4&per_page=20",
			strings.NewReader(""),
		),
	)

	fmt.Println(recorder.Header().Get("Link"))

	// Output:
	// page 4
	// perPage 20
	// offset 60
	// limit 80
	// <http://host.com/page%23anchor?custom=1&page=4&per_page=20> rel="next", <http://host.com/page%23anchor?custom=1&page=4&per_page=20> rel="prev", <http://host.com/page%23anchor?custom=1&page=4&per_page=20> rel="first", <http://host.com/page%23anchor?custom=1&page=4&per_page=20> rel="last"
}

func TestPaginator_Middleware(t *testing.T) {
	tests := []struct {
		name               string
		config             Config
		target             string
		expectedPagination *Pagination
	}{
		{
			name:               "Default config without params",
			config:             Config{},
			target:             "http://host.com/page",
			expectedPagination: &Pagination{page: 1, perPage: 100},
		},
		{
			name:               "Default config with params",
			config:             Config{},
			target:             "http://host.com/page?page=2&per_page=15",
			expectedPagination: &Pagination{page: 2, perPage: 15},
		},
		{
			name:               "Custom config without params",
			config:             Config{KeyPage: "customPage", KeyPerPage: "customPerPage", Page: 2, PerPage: 20},
			target:             "http://host.com/page?page=3&per_page=15",
			expectedPagination: &Pagination{page: 2, perPage: 20},
		},
		{
			name:               "Custom config with params",
			config:             Config{KeyPage: "customPage", KeyPerPage: "customPerPage", Page: 2, PerPage: 20},
			target:             "http://host.com/page?page=2&per_page=15&customPage=3&customPerPage=30",
			expectedPagination: &Pagination{page: 3, perPage: 30},
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
			paginator.Middleware(handler).ServeHTTP(
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodGet, test.target, strings.NewReader("")),
			)
		})
	}
}
