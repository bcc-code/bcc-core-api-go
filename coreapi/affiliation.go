package coreapi

import (
	"context"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
)

var AffiliationPath = "/affiliations"

type AffiliationManager managerBase

func (m *AffiliationManager) Get(ctx context.Context, uid strfmt.UUID, opts ...RequestOption) (models.WrappedAffiliation, error) {
	var res models.WrappedAffiliation
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(AffiliationPath, string(uid)), nil, &res, opts...)
	return res, err
}

func (m *AffiliationManager) Find(ctx context.Context, opts ...RequestOption) (models.WrappedWithMetaArrayAffiliation, error) {
	var res models.WrappedWithMetaArrayAffiliation
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(AffiliationPath), nil, &res, opts...)
	return res, err
}
