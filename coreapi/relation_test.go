package coreapi

import (
	"context"
	"testing"
	"time"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestGetRelation(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead, ScopePersonRelationsRead)

	res, err := c.Relation.Get(context.Background(), TestRelationUID)

	assert.NoError(t, err)
	assert.Equal(t, TestRelationUID, *res.Data.UID)
}

func TestFindRelation(t *testing.T) {
	c := getClientForTests(t, ScopePersonRelationsRead)

	res, err := c.Relation.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}

func TestCreateAndUpdateRelation(t *testing.T) {
	c := getClientForTests(t, ScopePersonRelationsWrite)

	body := models.RelationWrite{
		OriginUID:     lo.ToPtr(TestPersonUID),
		TargetUID:     lo.ToPtr(TestPerson2UID),
		Type:          models.RelationTypeChildOf,
		GrantToOrigin: models.GrantTypeDefault,
		GrantToTarget: models.GrantTypeDefault,
		ValidFrom:     lo.ToPtr(strfmt.DateTime(time.Now().Truncate(time.Millisecond))),
	}

	res, err := c.Relation.Create(context.Background(), body)
	assert.NoError(t, err)

	t.Run("CheckCreate", func(t *testing.T) {
		assert.NotEmpty(t, res.Data.UID)

		assert.Equal(t, body.OriginUID, res.Data.OriginUID)
		assert.Equal(t, body.TargetUID, res.Data.TargetUID)
		assert.Equal(t, body.Type, res.Data.Type)
		assert.True(t, body.ValidFrom.Equal(*res.Data.ValidFrom))
		assert.Equal(t, body.ValidTo, res.Data.ValidTo)

	})

	t.Run("CheckUpdate", func(t *testing.T) {
		body.ValidTo = lo.ToPtr(strfmt.DateTime(time.Now().Truncate(time.Millisecond)))
		res, err := c.Relation.Update(context.Background(), *res.Data.UID, body)

		assert.NoError(t, err)
		assert.NotEmpty(t, res.Data.UID)

		assert.Equal(t, body.OriginUID, res.Data.OriginUID)
		assert.Equal(t, body.TargetUID, res.Data.TargetUID)
		assert.Equal(t, body.Type, res.Data.Type)
		assert.True(t, body.ValidFrom.Equal(*res.Data.ValidFrom))
		assert.True(t, body.ValidTo.Equal(*res.Data.ValidTo))

	})
}
