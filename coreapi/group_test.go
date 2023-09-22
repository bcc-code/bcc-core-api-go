package coreapi

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

const TestGroupUID strfmt.UUID = "e88e4a8d-558a-4dfd-8e9e-389522ade4e0"
const TestGroupMemberUID strfmt.UUID = "603ee270-ca1f-41da-9349-944b0fa53a64"

func TestGetGroup(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopeGroupsRead)

	res, err := c.Group.Get(context.Background(), TestGroupUID)

	assert.NoError(t, err)
	assert.Equal(t, TestGroupUID, *res.Data.UID)
}

func TestFindGroup(t *testing.T) {
	c := GetTestClient(ScopeGroupsRead)

	res, err := c.Group.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}

func TestGetGroupMember(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopeGroupsRead)

	res, err := c.Group.GetMember(context.Background(), TestGroupUID, TestGroupMemberUID)

	assert.NoError(t, err)
	assert.Equal(t, TestGroupMemberUID, *res.Data.UID)
}

func TestGetFindMembers(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopeGroupsRead)

	res, err := c.Group.FindMembers(context.Background(), TestGroupUID, Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
