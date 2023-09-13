package coreapi

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Relation struct {
	Uid             uuid.UUID `json:"uid"`
	LastChangedDate time.Time `json:"lastChangedDate"`
	Origin          *Person   `json:"origin"`
	Target          *Person   `json:"target"`

	OriginUid uuid.UUID `json:"originUid"`
	TargetUid uuid.UUID `json:"targetUid"`

	// Type of the relation, defined as {origin} is {type} (of) {target}
	Type RelationType `json:"type"`

	// Permission of origin on target
	GrantToTarget Grant `json:"grantToTarget"`
	// Permission of target on origin
	GrantToOrigin Grant `json:"grantToOrigin"`

	ValidFrom time.Time  `json:"validFrom"`
	ValidTo   *time.Time `json:"validTo"`
}

// ENUM(ChildOf, SpouseOf, LegalDependentOf, FosterChildOf, ContactDependentOf)
type RelationType string

const (
	RelationTypeChildOf            RelationType = "ChildOf"
	RelationTypeSpouseOf           RelationType = "SpouseOf"
	RelationTypeLegalDependentOf   RelationType = "LegalDependentOf"
	RelationTypeFosterChildOf      RelationType = "FosterChildOf"
	RelationTypeContactDependentOf RelationType = "ContactDependentOf"
)

type Grant string

const (
	GrantDefault      Grant = "Default"
	GrantView         Grant = "View"
	GrantAdministrate Grant = "Administrate"
	GrantRepresent    Grant = "Represent"
	GrantNone         Grant = "None"
)

var RelationPath = "/relations"

type RelationManager managerBase

func (m *RelationManager) Get(ctx context.Context, uid uuid.UUID, opts ...RequestOption) (Response[Relation], error) {
	var res Response[Relation]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(RelationPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *RelationManager) Find(ctx context.Context, opts ...RequestOption) (ResponseWithMeta[[]Relation], error) {
	var res ResponseWithMeta[[]Relation]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(RelationPath), nil, &res, opts...)
	return res, err
}
