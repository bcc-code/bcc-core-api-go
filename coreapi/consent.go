package coreapi

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Consent struct {
	Uid             uuid.UUID `json:"uid"`
	LastChangedDate time.Time `json:"lastChangedDate"`
	Person          *Person   `json:"person"`
	Org             *Org      `json:"org"`

	PersonUid uuid.UUID      `json:"personUid"`
	OrgUid    uuid.UUID      `json:"orgUid"`
	Purpose   ConsentPurpose `json:"purpose"`
	ValidFrom time.Time      `json:"validFrom"`
	ValidTo   *time.Time     `json:"validTo"`
}

// ENUM(DataSharing, Tracking)
type ConsentPurpose string

var ConsentPath = "/consents"

type ConsentManager managerBase

func (m *ConsentManager) Get(ctx context.Context, uid uuid.UUID, opts ...RequestOption) (Response[Consent], error) {
	var res Response[Consent]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(ConsentPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *ConsentManager) Find(ctx context.Context, opts ...RequestOption) (ResponseWithMeta[[]Consent], error) {
	var res ResponseWithMeta[[]Consent]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(ConsentPath), nil, &res, opts...)
	return res, err
}
