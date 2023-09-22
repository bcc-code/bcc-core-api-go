package coreapi

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

const TestRoleAssignmentUID strfmt.UUID = "fff50a49-3413-4d1e-a99d-da66a45b7cd5"

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
