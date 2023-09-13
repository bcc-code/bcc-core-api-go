package coreapi

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var TestRoleAssignmentUid = uuid.MustParse("fff50a49-3413-4d1e-a99d-da66a45b7cd5")

func TestGetRoleAssignment(t *testing.T) {
	c := GetTestClient(PersonsRead, PersonRoleAssignmentsRead)

	res, err := c.RoleAssignment.Get(context.Background(), TestRoleAssignmentUid)

	assert.NoError(t, err)
	assert.Equal(t, TestRoleAssignmentUid, res.Data.Uid)
}

func TestFindRoleAssignment(t *testing.T) {
	c := GetTestClient(PersonRoleAssignmentsRead)

	res, err := c.RoleAssignment.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
