package bcccoreapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRole(t *testing.T) {
	c := getClientForTests(t, ScopeRolesRead)

	res, err := c.Role.Get(context.Background(), TestRoleUID)

	assert.NoError(t, err)
	assert.Equal(t, TestRoleUID, *res.Data.UID)
}

func TestFindRole(t *testing.T) {
	c := getClientForTests(t, ScopeRolesRead)

	res, err := c.Role.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
