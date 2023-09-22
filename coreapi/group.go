package coreapi

import (
	"context"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
)

var GroupPath = "/groups"

type GroupManager managerBase

func (m *GroupManager) Get(ctx context.Context, uid strfmt.UUID, opts ...RequestOption) (models.WrappedGroup, error) {
	var res models.WrappedGroup
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(GroupPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *GroupManager) Find(ctx context.Context, opts ...RequestOption) (models.WrappedWithMetaArrayGroup, error) {
	var res models.WrappedWithMetaArrayGroup
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(GroupPath), nil, &res, opts...)
	return res, err
}

var MemberPath = "/members"

func (m *GroupManager) GetMember(ctx context.Context, groupUID strfmt.UUID, memberUID strfmt.UUID, opts ...RequestOption) (models.WrappedGroupMember, error) {
	var res models.WrappedGroupMember
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(GroupPath, groupUID.String(), MemberPath, memberUID.String()), nil, &res, opts...)
	return res, err
}

func (m *GroupManager) FindMembers(ctx context.Context, groupUID strfmt.UUID, opts ...RequestOption) (models.WrappedWithMetaArrayGroupMember, error) {
	var res models.WrappedWithMetaArrayGroupMember
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(GroupPath, groupUID.String(), MemberPath), nil, &res, opts...)
	return res, err
}
