package coreapi

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Country struct {
	Uid             uuid.UUID `json:"uid"`
	LastChangedDate time.Time `json:"lastChangedDate"`

	Iso2Code   string `json:"iso2Code"`
	NameEn     string `json:"nameEn"`
	NameNative string `json:"nameNative"`
	NameNo     string `json:"nameNo"`
}

var CountryPath = "/countries"

type CountryManager managerBase

func (m *CountryManager) Get(ctx context.Context, uid uuid.UUID, opts ...RequestOption) (Response[Country], error) {
	var res Response[Country]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(CountryPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *CountryManager) Find(ctx context.Context, opts ...RequestOption) (ResponseWithMeta[[]Country], error) {
	var res ResponseWithMeta[[]Country]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(CountryPath), nil, &res, opts...)
	return res, err
}
