//go:generate go-enum --marshal

package coreapi

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Org struct {
	Uid             uuid.UUID       `json:"uid"`
	OrgID           int             `json:"orgID"`
	LastChangedDate time.Time       `json:"lastChangedDate"`
	Name            string          `json:"name"`
	DistrictName    string          `json:"districtName"`
	Type            OrgType         `json:"type"`
	ActiveStatus    OrgActiveStatus `json:"activeStatus"`
	PostalAddress   Address         `json:"postalAddress"`
	BillingAddress  Address         `json:"billingAddress"`
	VisitingAddress Address         `json:"visitingAddress"`
}

// ENUM(Church, Club, Org)
type OrgType string

// ENUM(Active, Inactive)
type OrgActiveStatus string

var OrgPath = "/v2/orgs"

type OrgManager managerBase

func (m *OrgManager) Get(ctx context.Context, uid uuid.UUID, opts ...RequestOption) (Response[Org], error) {
	var res Response[Org]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(OrgPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *OrgManager) Find(ctx context.Context, opts ...RequestOption) (ResponseWithMeta[[]Org], error) {
	var res ResponseWithMeta[[]Org]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(OrgPath), nil, &res, opts...)
	return res, err
}
