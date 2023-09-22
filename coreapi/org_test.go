package coreapi

import (
	"context"
	"testing"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestGetOrg(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead, ScopeOrgsRead)

	res, err := c.Org.Get(context.Background(), TestOrgUID)

	assert.NoError(t, err)
	assert.Equal(t, TestOrgUID, *res.Data.UID)
}

func TestFindOrg(t *testing.T) {
	c := getClientForTests(t, ScopeOrgsRead)

	res, err := c.Org.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}

func TestCreateAndUpdateOrg(t *testing.T) {
	c := getClientForTests(t, ScopeOrgsWrite)

	body := models.OrgWrite{
		Name:         lo.ToPtr("BCC Svartskog"),
		DistrictName: lo.ToPtr("Svartskog"),
		ActiveStatus: models.OrgActiveStatusActive,
		BillingAddress: &models.Address{
			Address1: "Svartskog 1",
		},
		Type: models.OrgTypeChurch,
	}

	res, err := c.Org.Create(context.Background(), body)
	assert.NoError(t, err)

	t.Run("CheckCreate", func(t *testing.T) {
		assert.NotEmpty(t, res.Data.UID)
		assert.NotZero(t, res.Data.OrgID)

		assert.Equal(t, body.Name, res.Data.Name)
		assert.Equal(t, body.DistrictName, res.Data.DistrictName)
		assert.Equal(t, body.ActiveStatus, res.Data.ActiveStatus)
		assert.Equal(t, body.BillingAddress, res.Data.BillingAddress)
		assert.Equal(t, body.Type, res.Data.Type)

	})

	t.Run("CheckUpdate", func(t *testing.T) {
		body.Name = lo.ToPtr("BCC New Svartskog")
		res, err := c.Org.Update(context.Background(), *res.Data.UID, body)

		assert.NoError(t, err)

		assert.Equal(t, body.Name, res.Data.Name)
		assert.Equal(t, body.DistrictName, res.Data.DistrictName)
		assert.Equal(t, body.ActiveStatus, res.Data.ActiveStatus)
		assert.Equal(t, body.BillingAddress, res.Data.BillingAddress)
		assert.Equal(t, body.Type, res.Data.Type)
	})
}
