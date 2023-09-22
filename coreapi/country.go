package coreapi

import (
	"context"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
)

var CountryPath = "/countries"

type CountryManager managerBase

func (m *CountryManager) Get(ctx context.Context, uid strfmt.UUID, opts ...RequestOption) (models.WrappedCountry, error) {
	var res models.WrappedCountry
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(CountryPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *CountryManager) Find(ctx context.Context, opts ...RequestOption) (models.WrappedWithMetaArrayCountry, error) {
	var res models.WrappedWithMetaArrayCountry
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(CountryPath), nil, &res, opts...)
	return res, err
}
