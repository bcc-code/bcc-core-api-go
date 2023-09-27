package bcccoreapi

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReuqestOptions(t *testing.T) {
	testCases := []struct {
		Opts []RequestOption
		Res  string
	}{
		{[]RequestOption{}, ""},
		{[]RequestOption{Limit(100)}, "limit=100"},
		{[]RequestOption{Limit(20), Skip(100)}, "limit=20&skip=100"},
		{[]RequestOption{Page(5), Limit(10)}, "limit=10&page=5"},
		{[]RequestOption{Fields("*", "test")}, fmt.Sprintf("fields=%s", url.QueryEscape("*,test"))},
		{[]RequestOption{Sort("-test")}, "sort=-test"},
		{[]RequestOption{Filter(`{"displayName": {"_eq": "Test"}}`)},
			fmt.Sprintf("filter=%s", url.QueryEscape(`{"displayName": {"_eq": "Test"}}`))},
	}

	for _, tc := range testCases {
		t.Run(tc.Res, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "example", nil)
			for _, opt := range tc.Opts {
				opt(req)
			}
			assert.Equal(t, tc.Res, req.URL.RawQuery)
		})
	}
}
