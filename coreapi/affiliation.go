//go:generate go-enum --marshal

package coreapi

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Affiliation struct {
	Uid             uuid.UUID `json:"uid"`
	LastChangedDate time.Time `json:"lastChangedDate"`
	Org             *Org      `json:"org"`
	Person          *Person   `json:"person"`

	PersonUid uuid.UUID       `json:"personUid"`
	OrgUid    uuid.UUID       `json:"orgUid"`
	Type      AffiliationType `json:"type"`
	ValidFrom time.Time       `json:"validFrom"`
	ValidTo   *time.Time      `json:"validTo"`
}

// ENUM(Member, Affiliate, Participant)
type AffiliationType string

var AffiliationPath = "/affiliations"

type AffiliationManager managerBase

func (m *AffiliationManager) Get(ctx context.Context, uid uuid.UUID, opts ...RequestOption) (Response[Affiliation], error) {
	var res Response[Affiliation]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(AffiliationPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *AffiliationManager) Find(ctx context.Context, opts ...RequestOption) (ResponseWithMeta[[]Affiliation], error) {
	var res ResponseWithMeta[[]Affiliation]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(AffiliationPath), nil, &res, opts...)
	return res, err
}
