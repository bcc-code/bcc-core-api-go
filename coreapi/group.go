package coreapi

import (
	"context"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
)

var GroupPath = "/groups"

type GroupClient struct {
	genericClient[models.Group, models.GroupWrite]
}

const MemberPath = "/members"

func (m *GroupClient) GetMember(ctx context.Context, groupUID strfmt.UUID, memberUID strfmt.UUID, opts ...RequestOption) (models.Wrapped[models.GroupMember], error) {
	var res models.Wrapped[models.GroupMember]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(GroupPath, groupUID.String(), MemberPath, memberUID.String()), nil, &res, opts...)
	return res, err
}

func (m *GroupClient) FindMembers(ctx context.Context, groupUID strfmt.UUID, opts ...RequestOption) (models.WrappedWithMeta[[]models.GroupMember], error) {
	var res models.WrappedWithMeta[[]models.GroupMember]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(GroupPath, groupUID.String(), MemberPath), nil, &res, opts...)
	return res, err
}

func (m *GroupClient) AddMember(ctx context.Context, groupUID strfmt.UUID, member models.GroupMemberWrite, opts ...RequestOption) (models.Wrapped[models.GroupMember], error) {
	var res models.Wrapped[models.GroupMember]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(GroupPath, groupUID.String(), MemberPath), member, &res, opts...)
	return res, err
}

func (m *GroupClient) RemoveMember(ctx context.Context, groupUID strfmt.UUID, memberUID strfmt.UUID, opts ...RequestOption) (models.Wrapped[models.GroupMember], error) {
	var res models.Wrapped[models.GroupMember]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(GroupPath, groupUID.String(), MemberPath, memberUID.String()), nil, &res, opts...)
	return res, err
}
