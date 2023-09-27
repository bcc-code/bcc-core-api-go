package bcccoreapi

import (
	"context"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
)

type groupClient struct {
	genericClient[models.Group, models.GroupWrite]
}

func (m *groupClient) GetMember(ctx context.Context, groupUID strfmt.UUID, memberUID strfmt.UUID, opts ...RequestOption) (models.Wrapped[models.GroupMember], error) {
	var res models.Wrapped[models.GroupMember]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(groupsPath, groupUID.String(), groupMembersPath, memberUID.String()), nil, &res, opts...)
	return res, err
}

func (m *groupClient) FindMembers(ctx context.Context, groupUID strfmt.UUID, opts ...RequestOption) (models.WrappedWithMeta[[]models.GroupMember], error) {
	var res models.WrappedWithMeta[[]models.GroupMember]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(groupsPath, groupUID.String(), groupMembersPath), nil, &res, opts...)
	return res, err
}

func (m *groupClient) AddMember(ctx context.Context, groupUID strfmt.UUID, member models.GroupMemberWrite, opts ...RequestOption) (models.Wrapped[models.GroupMember], error) {
	var res models.Wrapped[models.GroupMember]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(groupsPath, groupUID.String(), groupMembersPath), member, &res, opts...)
	return res, err
}

func (m *groupClient) RemoveMember(ctx context.Context, groupUID strfmt.UUID, memberUID strfmt.UUID, opts ...RequestOption) (models.Wrapped[models.GroupMember], error) {
	var res models.Wrapped[models.GroupMember]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(groupsPath, groupUID.String(), groupMembersPath, memberUID.String()), nil, &res, opts...)
	return res, err
}
