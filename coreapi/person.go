package coreapi

import (
	"context"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
)

var PersonPath = "/v2/persons"

type PersonManager managerBase

func (m *PersonManager) Get(ctx context.Context, uid strfmt.UUID, opts ...RequestOption) (models.WrappedPerson, error) {
	var res models.WrappedPerson
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(PersonPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *PersonManager) Find(ctx context.Context, opts ...RequestOption) (models.WrappedWithMetaArrayPerson, error) {
	var res models.WrappedWithMetaArrayPerson
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(PersonPath), nil, &res, opts...)
	return res, err
}
