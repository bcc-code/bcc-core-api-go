package coreapi

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var TestOrgUid = uuid.MustParse("9e0b8d03-9799-416d-a70b-6da9f2e1ab48")

func TestGetOrg(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopeOrgsRead)

	res, err := c.Org.Get(context.Background(), TestOrgUid)

	assert.NoError(t, err)
	assert.Equal(t, TestOrgUid, res.Data.Uid)
}

func TestFindOrg(t *testing.T) {
	c := GetTestClient(ScopeOrgsRead)

	res, err := c.Org.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
