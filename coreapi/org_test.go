package coreapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrg(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopeOrgsRead)

	res, err := c.Org.Get(context.Background(), TestOrgUID)

	assert.NoError(t, err)
	assert.Equal(t, TestOrgUID, *res.Data.UID)
}

func TestFindOrg(t *testing.T) {
	c := GetTestClient(ScopeOrgsRead)

	res, err := c.Org.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
