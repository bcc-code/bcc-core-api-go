package coreapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRoleAssignment(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopePersonRoleAssignmentsRead)

	res, err := c.RoleAssignment.Get(context.Background(), TestRoleAssignmentUID)

	assert.NoError(t, err)
	assert.Equal(t, TestRoleAssignmentUID, *res.Data.UID)
}

func TestFindRoleAssignment(t *testing.T) {
	c := GetTestClient(ScopePersonRoleAssignmentsRead)

	res, err := c.RoleAssignment.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
