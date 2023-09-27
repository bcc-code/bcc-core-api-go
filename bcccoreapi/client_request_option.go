package bcccoreapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type RequestOption func(*http.Request)

func Limit(n int) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("limit", fmt.Sprint(n))
		req.URL.RawQuery = q.Encode()
	}
}

func Skip(n int) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("skip", fmt.Sprint(n))
		req.URL.RawQuery = q.Encode()
	}
}

func Page(n int) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("page", fmt.Sprint(n))
		req.URL.RawQuery = q.Encode()
	}
}

func Fields(fields ...string) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("fields", strings.Join(fields, ","))
		req.URL.RawQuery = q.Encode()
	}
}

func Sort(sortFields ...string) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("sort", strings.Join(sortFields, ","))
		req.URL.RawQuery = q.Encode()
	}
}

func Filter(filter string) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("filter", filter)
		req.URL.RawQuery = q.Encode()
	}
}

type QueryOpts struct {
	IncludePersonsWithoutChurchAffiliation bool     `json:"includePersonsWithoutChurchAffiliation"`
	IncludeInactive                        []string `json:"includeInactive"`
	OrgUIDs                                []uuid.UUID
}

func QueryOptions(q QueryOpts) RequestOption {
	queryOptionsJson, _ := json.Marshal(q)
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("queryOptions", string(queryOptionsJson))
		req.URL.RawQuery = q.Encode()
	}
}

func Search(s string) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("search", s)
		req.URL.RawQuery = q.Encode()
	}
}
