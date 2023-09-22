package coreapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
