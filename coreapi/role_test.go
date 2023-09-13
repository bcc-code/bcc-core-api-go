package coreapi

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var TestRoleUid = uuid.MustParse("47b2dd20-d31c-4b07-8837-647ea8ee0f7d")

func TestGetRole(t *testing.T) {
	c := GetTestClient(RolesRead)

	res, err := c.Role.Get(context.Background(), TestRoleUid)

	assert.NoError(t, err)
	assert.Equal(t, TestRoleUid, res.Data.Uid)
}

func TestFindRole(t *testing.T) {
	c := GetTestClient(RolesRead)

	res, err := c.Role.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
