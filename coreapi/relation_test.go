package coreapi

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

const TestRelationUID strfmt.UUID = "3a0c7609-bd9e-4f1e-a5b2-e8471ebc7b92"

func TestGetRelation(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopePersonRelationsRead)

	res, err := c.Relation.Get(context.Background(), TestRelationUID)

	assert.NoError(t, err)
	assert.Equal(t, TestRelationUID, *res.Data.UID)
}

func TestFindRelation(t *testing.T) {
	c := GetTestClient(ScopePersonRelationsRead)

	res, err := c.Relation.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
