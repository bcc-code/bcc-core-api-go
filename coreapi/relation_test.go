package coreapi

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var TestRelationUid = uuid.MustParse("3a0c7609-bd9e-4f1e-a5b2-e8471ebc7b92")

func TestGetRelation(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopePersonRelationsRead)

	res, err := c.Relation.Get(context.Background(), TestRelationUid)

	assert.NoError(t, err)
	assert.Equal(t, TestRelationUid, res.Data.Uid)
}

func TestFindRelation(t *testing.T) {
	c := GetTestClient(ScopePersonRelationsRead)

	res, err := c.Relation.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
