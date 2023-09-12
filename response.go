package coreapi

type Response[T any] struct {
	Data T `json:"data"`
}

type ResponseWithMeta[T any] struct {
	Data T    `json:"data"`
	Meta Meta `json:"meta"`
}

type Meta struct {
	Limit   int `json:"limit"`
	Skipped int `json:"skipped"`
	Total   int `json:"total"`
}
