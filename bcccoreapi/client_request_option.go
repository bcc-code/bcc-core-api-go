package bcccoreapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type RequestOption func(*http.Request)

// Limit will restrict the number of returned results to n
// See more here https://developer.bcc.no/bcc-core-api/crud/query-parameters.html#limit
// Only works for Find endpoints
func Limit(n int) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("limit", fmt.Sprint(n))
		req.URL.RawQuery = q.Encode()
	}
}

// Skip will skip the first n rows from the results
// See more here https://developer.bcc.no/bcc-core-api/crud/query-parameters.html#skip
// Only works for Find endpoints
func Skip(n int) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("skip", fmt.Sprint(n))
		req.URL.RawQuery = q.Encode()
	}
}

// Skip will skip the first (n - 1) * Limit rows from the results
// See more here https://developer.bcc.no/bcc-core-api/crud/query-parameters.html#page
// Only works for Find endpoints
func Page(n int) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("page", fmt.Sprint(n))
		req.URL.RawQuery = q.Encode()
	}
}

// Fields allows to specify the fields returned by the SDK
// See more here https://developer.bcc.no/bcc-core-api/crud/query-parameters.html#fields
// Works for Find and Get endpoints
func Fields(fields ...string) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("fields", strings.Join(fields, ","))
		req.URL.RawQuery = q.Encode()
	}
}

// Sorts the result by a given field
// See more here https://developer.bcc.no/bcc-core-api/crud/query-parameters.html#sort
// Only works for Find endpoints
func Sort(sortFields ...string) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("sort", strings.Join(sortFields, ","))
		req.URL.RawQuery = q.Encode()
	}
}

// Filter the result based on the provided condition
// See more here https://developer.bcc.no/bcc-core-api/crud/query-parameters.html#filter
// Only works for Find endpoints
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

// Sets additional query options
// See more here https://developer.bcc.no/bcc-core-api/crud/query-parameters.html#options
// Works for Find and Get endpoints
func QueryOptions(q QueryOpts) RequestOption {
	queryOptionsJson, _ := json.Marshal(q)
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("queryOptions", string(queryOptionsJson))
		req.URL.RawQuery = q.Encode()
	}
}

// Performs a fuzzy search
// See more here https://developer.bcc.no/bcc-core-api/crud/query-parameters.html#search
// Only works for Find Person endpoint
func Search(s string) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set("search", s)
		req.URL.RawQuery = q.Encode()
	}
}
