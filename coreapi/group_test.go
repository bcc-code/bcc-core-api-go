package coreapi

import (
	"context"
	"testing"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestGetGroup(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead, ScopeGroupsRead)

	res, err := c.Group.Get(context.Background(), TestGroupUID)

	assert.NoError(t, err)
	assert.Equal(t, TestGroupUID, *res.Data.UID)
}

func TestFindGroup(t *testing.T) {
	c := getClientForTests(t, ScopeGroupsRead)

	res, err := c.Group.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}

func TestGetGroupMember(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead, ScopeGroupsRead)

	res, err := c.Group.GetMember(context.Background(), TestGroupUID, TestGroupMemberUID)

	assert.NoError(t, err)
	assert.Equal(t, TestGroupMemberUID, *res.Data.UID)
}

func TestGetFindMembers(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead, ScopeGroupsRead)

	res, err := c.Group.FindMembers(context.Background(), TestGroupUID, Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}

func TestCreateAndUpdateGroup(t *testing.T) {
	c := getClientForTests(t, ScopeGroupsWrite)

	body := models.GroupWrite{
		OrgUID: TestOrgUID,
		Name:   lo.ToPtr("test-group"),
		Rule:   `{"displayName": {"_gt": "A"}}`,
		Tags:   []string{"testing"},
		Type:   models.GroupTypeDynamic.Pointer(),
	}

	res, err := c.Group.Create(context.Background(), body)
	assert.NoError(t, err)

	t.Run("CheckCreate", func(t *testing.T) {
		assert.NotEmpty(t, res.Data.UID)

		assert.Equal(t, body.OrgUID, res.Data.OrgUID)
		assert.Equal(t, body.Rule, res.Data.Rule)
		assert.Equal(t, body.Tags, res.Data.Tags)
		assert.Equal(t, body.Type, res.Data.Type)

	})

	t.Run("CheckUpdate", func(t *testing.T) {
		body.Tags = append(body.Tags, "updatedTag")
		res, err := c.Group.Update(context.Background(), *res.Data.UID, body)

		assert.NoError(t, err)
		assert.NotEmpty(t, res.Data.UID)

		assert.Equal(t, body.OrgUID, res.Data.OrgUID)
		assert.Equal(t, body.Rule, res.Data.Rule)
		assert.Equal(t, body.Tags, res.Data.Tags)
		assert.Equal(t, body.Type, res.Data.Type)
	})
}
