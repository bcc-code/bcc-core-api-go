package coreapi

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var TestGroupUid = uuid.MustParse("e88e4a8d-558a-4dfd-8e9e-389522ade4e0")
var TestGroupMemberUid = uuid.MustParse("603ee270-ca1f-41da-9349-944b0fa53a64")

func TestGetGroup(t *testing.T) {
	c := GetTestClient(PersonsRead, GroupsRead)

	res, err := c.Group.Get(context.Background(), TestGroupUid)

	assert.NoError(t, err)
	assert.Equal(t, TestGroupUid, res.Data.Uid)
}

func TestFindGroup(t *testing.T) {
	c := GetTestClient(GroupsRead)

	res, err := c.Group.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}

func TestGetGroupMember(t *testing.T) {
	c := GetTestClient(PersonsRead, GroupsRead)

	res, err := c.Group.GetMember(context.Background(), TestGroupUid, TestGroupMemberUid)

	assert.NoError(t, err)
	assert.Equal(t, TestGroupMemberUid, res.Data.Uid)
}

func TestGetFindMembers(t *testing.T) {
	c := GetTestClient(PersonsRead, GroupsRead)

	res, err := c.Group.FindMembers(context.Background(), TestGroupUid, Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
