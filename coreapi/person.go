package coreapi

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type Person struct {
	Uid         uuid.UUID `json:"uid"`
	DisplayName string    `json:"displayName"`
}

var PersonPath = "/v2/persons"

type PersonManager managerBase

func (m *PersonManager) Get(ctx context.Context, uid uuid.UUID, opts ...RequestOption) (Person, error) {
	var p Response[Person]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(PersonPath, uid.String()), nil, &p, opts...)
	return p.Data, err
}

func (m *PersonManager) Find(ctx context.Context, opts ...RequestOption) (ResponseWithMeta[[]Person], error) {
	var p ResponseWithMeta[[]Person]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(PersonPath), nil, &p, opts...)
	return p, err
}
