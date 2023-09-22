package coreapi

import (
	"context"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
)

var RolePath = "/roles"

type RoleManager managerBase

func (m *RoleManager) Get(ctx context.Context, uid strfmt.UUID, opts ...RequestOption) (models.WrappedRole, error) {
	var res models.WrappedRole
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(RolePath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *RoleManager) Find(ctx context.Context, opts ...RequestOption) (models.WrappedWithMetaArrayRole, error) {
	var res models.WrappedWithMetaArrayRole
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(RolePath), nil, &res, opts...)
	return res, err
}
