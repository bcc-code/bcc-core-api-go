package coreapi

import (
	"context"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
)

var RoleAssignmentPath = "/roleAssignments"

type RoleAssignmentManager managerBase

func (m *RoleAssignmentManager) Get(ctx context.Context, uid strfmt.UUID, opts ...RequestOption) (models.WrappedRoleAssignment, error) {
	var res models.WrappedRoleAssignment
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(RoleAssignmentPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *RoleAssignmentManager) Find(ctx context.Context, opts ...RequestOption) (models.WrappedWithMetaArrayRoleAssignment, error) {
	var res models.WrappedWithMetaArrayRoleAssignment
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(RoleAssignmentPath), nil, &res, opts...)
	return res, err
}
