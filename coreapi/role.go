//go:generate go-enum --marshal

package coreapi

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Role struct {
	Uid             uuid.UUID `json:"uid"`
	LastChangedDate time.Time `json:"lastChangedDate"`

	Name  string    `json:"name"`
	Scope RoleScope `json:"scope"`
}

// ENUM(Global, Org)
type RoleScope string //@name RoleScope

var RolePath = "/roles"

type RoleManager managerBase

func (m *RoleManager) Get(ctx context.Context, uid uuid.UUID, opts ...RequestOption) (Response[Role], error) {
	var res Response[Role]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(RolePath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *RoleManager) Find(ctx context.Context, opts ...RequestOption) (ResponseWithMeta[[]Role], error) {
	var res ResponseWithMeta[[]Role]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(RolePath), nil, &res, opts...)
	return res, err
}
