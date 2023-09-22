package coreapi

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

const TestOrgUID strfmt.UUID = "9e0b8d03-9799-416d-a70b-6da9f2e1ab48"

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
