package coreapi

import (
	"context"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
)

var RelationPath = "/relations"

type RelationManager managerBase

func (m *RelationManager) Get(ctx context.Context, uid strfmt.UUID, opts ...RequestOption) (models.WrappedRelation, error) {
	var res models.WrappedRelation
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(RelationPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *RelationManager) Find(ctx context.Context, opts ...RequestOption) (models.WrappedWithMetaArrayRelation, error) {
	var res models.WrappedWithMetaArrayRelation
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(RelationPath), nil, &res, opts...)
	return res, err
}
