package coreapi

import (
	"context"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
)

var ConsentPath = "/consents"

type ConsentManager managerBase

func (m *ConsentManager) Get(ctx context.Context, uid strfmt.UUID, opts ...RequestOption) (models.WrappedConsent, error) {
	var res models.WrappedConsent
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(ConsentPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *ConsentManager) Find(ctx context.Context, opts ...RequestOption) (models.WrappedWithMetaArrayConsent, error) {
	var res models.WrappedWithMetaArrayConsent
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(ConsentPath), nil, &res, opts...)
	return res, err
}
