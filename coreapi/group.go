package coreapi

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Group struct {
	Uid             uuid.UUID `json:"uid"`
	LastChangedDate time.Time `json:"lastChangedDate"`

	AppUid uuid.UUID `json:"appUid"`

	OrgUid *uuid.UUID `json:"orgUid"`
	Name   string     `json:"name"`
	Type   GroupType  `json:"type"`
	Rule   string     `json:"rule"`
	Tags   []string   `json:"tags"`
}

// ENUM( Static, Dynamic )
type GroupType string // @name GroupType

type GroupMember struct {
	Uid             uuid.UUID `json:"uid"`
	LastChangedDate time.Time `json:"lastChangedDate"`

	Person    *Person   `json:"person"`
	GroupUid  uuid.UUID `json:"groupUid"`
	PersonUid uuid.UUID `json:"personUid"`
}

var GroupPath = "/groups"

type GroupManager managerBase

func (m *GroupManager) Get(ctx context.Context, uid uuid.UUID, opts ...RequestOption) (Response[Group], error) {
	var res Response[Group]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(GroupPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *GroupManager) Find(ctx context.Context, opts ...RequestOption) (ResponseWithMeta[[]Group], error) {
	var res ResponseWithMeta[[]Group]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(GroupPath), nil, &res, opts...)
	return res, err
}

var MemberPath = "/members"

func (m *GroupManager) GetMember(ctx context.Context, groupUid uuid.UUID, memberUid uuid.UUID, opts ...RequestOption) (Response[GroupMember], error) {
	var res Response[GroupMember]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(GroupPath, groupUid.String(), MemberPath, memberUid.String()), nil, &res, opts...)
	return res, err
}

func (m *GroupManager) FindMembers(ctx context.Context, groupUid uuid.UUID, opts ...RequestOption) (ResponseWithMeta[[]GroupMember], error) {
	var res ResponseWithMeta[[]GroupMember]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(GroupPath, groupUid.String(), MemberPath), nil, &res, opts...)
	return res, err
}
