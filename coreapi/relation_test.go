package coreapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
