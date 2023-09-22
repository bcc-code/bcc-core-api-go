package coreapi

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

const TestRoleUID strfmt.UUID = "47b2dd20-d31c-4b07-8837-647ea8ee0f7d"

func TestGetRole(t *testing.T) {
	c := GetTestClient(ScopeRolesRead)

	res, err := c.Role.Get(context.Background(), TestRoleUID)

	assert.NoError(t, err)
	assert.Equal(t, TestRoleUID, *res.Data.UID)
}

func TestFindRole(t *testing.T) {
	c := GetTestClient(ScopeRolesRead)

	res, err := c.Role.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
