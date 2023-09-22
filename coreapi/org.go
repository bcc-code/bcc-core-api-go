package coreapi

import (
	"context"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
)

var OrgPath = "/v2/orgs"

type OrgManager managerBase

func (m *OrgManager) Get(ctx context.Context, uid strfmt.UUID, opts ...RequestOption) (models.WrappedOrg, error) {
	var res models.WrappedOrg
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(OrgPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *OrgManager) Find(ctx context.Context, opts ...RequestOption) (models.WrappedWithMetaArrayOrg, error) {
	var res models.WrappedWithMetaArrayOrg
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(OrgPath), nil, &res, opts...)
	return res, err
}
