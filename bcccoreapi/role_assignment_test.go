package bcccoreapi

import (
	"context"
	"testing"
	"time"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestGetRoleAssignment(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead, ScopePersonRoleAssignmentsRead)

	res, err := c.RoleAssignment.Get(context.Background(), TestRoleAssignmentUID)

	assert.NoError(t, err)
	assert.Equal(t, TestRoleAssignmentUID, *res.Data.UID)
}

func TestFindRoleAssignment(t *testing.T) {
	c := getClientForTests(t, ScopePersonRoleAssignmentsRead)

	res, err := c.RoleAssignment.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}

func TestCreateAndUpdateRoleAssignment(t *testing.T) {
	c := getClientForTests(t, ScopePersonRoleAssignmentsWrite)

	body := models.RoleAssignmentWrite{
		PersonUID: lo.ToPtr(TestPersonUID),
		RoleUID:   lo.ToPtr(TestRoleUID),
		ValidFrom: lo.ToPtr(strfmt.DateTime(time.Now().Truncate(time.Millisecond))),
	}

	res, err := c.RoleAssignment.Create(context.Background(), body)
	assert.NoError(t, err)

	t.Run("CheckCreate", func(t *testing.T) {
		assert.NotEmpty(t, res.Data.UID)

		assert.Equal(t, body.PersonUID, res.Data.PersonUID)
		assert.Equal(t, body.RoleUID, res.Data.RoleUID)
		assert.True(t, body.ValidFrom.Equal(*res.Data.ValidFrom))
		assert.Equal(t, body.ValidTo, res.Data.ValidTo)

	})

	t.Run("CheckUpdate", func(t *testing.T) {
		body.ValidTo = lo.ToPtr(strfmt.DateTime(time.Now().Truncate(time.Millisecond)))
		res, err := c.RoleAssignment.Update(context.Background(), *res.Data.UID, body)

		assert.NoError(t, err)
		assert.NotEmpty(t, res.Data.UID)

		assert.Equal(t, body.PersonUID, res.Data.PersonUID)
		assert.Equal(t, body.RoleUID, res.Data.RoleUID)
		assert.True(t, body.ValidFrom.Equal(*res.Data.ValidFrom))
		assert.True(t, body.ValidTo.Equal(*res.Data.ValidTo))

	})
}
