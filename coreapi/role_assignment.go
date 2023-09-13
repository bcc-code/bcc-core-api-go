package coreapi

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type RoleAssignment struct {
	Uid             uuid.UUID `json:"uid"`
	LastChangedDate time.Time `json:"lastChangedDate"`
	Person          *Person   `json:"person"`
	Role            *Role     `json:"role"`
	Org             *Org      `json:"org"`

	PersonUid uuid.UUID  `json:"personUid"`
	RoleUid   uuid.UUID  `json:"roleUid"`
	OrgUid    *uuid.UUID `json:"orgUid"`
	ValidFrom time.Time  `json:"validFrom"`
	ValidTo   *time.Time `json:"validTo"`
}

var RoleAssignmentPath = "/roleAssignments"

type RoleAssignmentManager managerBase

func (m *RoleAssignmentManager) Get(ctx context.Context, uid uuid.UUID, opts ...RequestOption) (Response[RoleAssignment], error) {
	var res Response[RoleAssignment]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(RoleAssignmentPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *RoleAssignmentManager) Find(ctx context.Context, opts ...RequestOption) (ResponseWithMeta[[]RoleAssignment], error) {
	var res ResponseWithMeta[[]RoleAssignment]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(RoleAssignmentPath), nil, &res, opts...)
	return res, err
}
