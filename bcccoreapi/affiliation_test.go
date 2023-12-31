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

func TestGetAffiliation(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead, ScopePersonAffiliationsRead)

	res, err := c.Affiliation.Get(context.Background(), TestAffiliationUID, Fields("*", "person.*"))

	assert.NoError(t, err)
	assert.Equal(t, TestAffiliationUID, *res.Data.UID)
	assert.NotZero(t, TestPersonUID, *res.Data.Person)
}

func TestFindAffiliation(t *testing.T) {
	c := getClientForTests(t, ScopePersonAffiliationsRead)

	res, err := c.Affiliation.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}

func TestCreateAffiliation(t *testing.T) {
	c := getClientForTests(t, ScopePersonAffiliationsWrite)

	body := models.AffiliationWrite{
		OrgUID:    lo.ToPtr(TestOrgUID),
		PersonUID: lo.ToPtr(TestPersonUID),
		Type:      models.AffiliationTypeAffiliate.Pointer(),
		ValidFrom: lo.ToPtr(strfmt.DateTime(time.Now().Truncate(time.Millisecond))),
	}

	res, err := c.Affiliation.Create(context.Background(), body)

	assert.NoError(t, err)
	assert.NotEmpty(t, res.Data.UID)

	assert.Equal(t, body.OrgUID, res.Data.OrgUID)
	assert.Equal(t, body.PersonUID, res.Data.PersonUID)
	assert.Equal(t, body.Type, res.Data.Type)

	assert.True(t, body.ValidFrom.Equal(*res.Data.ValidFrom))
	assert.Nil(t, body.ValidTo)
}

func TestUpdateAffiliation(t *testing.T) {
	c := getClientForTests(t, ScopePersonAffiliationsWrite, ScopePersonAffiliationsRead)

	currentAff, err := c.Affiliation.Get(context.Background(), TestAffiliationUID)
	assert.NoError(t, err)

	body := models.AffiliationWrite{
		OrgUID:    currentAff.Data.OrgUID,
		PersonUID: currentAff.Data.PersonUID,
		Type:      currentAff.Data.Type,
		ValidFrom: currentAff.Data.ValidFrom,
		ValidTo:   lo.ToPtr(strfmt.DateTime(time.Now().Truncate(time.Millisecond))),
	}

	res, err := c.Affiliation.Update(context.Background(), TestAffiliationUID, body)

	assert.NoError(t, err)
	assert.NotEmpty(t, res.Data.UID)

	assert.Equal(t, body.OrgUID, res.Data.OrgUID)
	assert.Equal(t, body.PersonUID, res.Data.PersonUID)
	assert.Equal(t, body.Type, res.Data.Type)

	assert.True(t, body.ValidFrom.Equal(*res.Data.ValidFrom))
	assert.True(t, body.ValidTo.Equal(*res.Data.ValidTo))
}
