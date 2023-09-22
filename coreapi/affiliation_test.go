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

func TestGetAffiliation(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopePersonAffiliationsRead)

	res, err := c.Affiliation.Get(context.Background(), TestAffiliationUID, Fields("*", "person.*"))

	assert.NoError(t, err)
	assert.Equal(t, TestAffiliationUID, *res.Data.UID)
	assert.NotZero(t, TestPersonUID, *res.Data.Person)
}

func TestFindAffiliation(t *testing.T) {
	c := GetTestClient(ScopePersonAffiliationsRead)

	res, err := c.Affiliation.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}

func TestCreateAffiliation(t *testing.T) {
	c := GetTestClient(ScopePersonAffiliationsWrite)

	body := models.AffiliationWrite{
		OrgUID:    lo.ToPtr(TestOrgUID),
		PersonUID: lo.ToPtr(TestPersonUID),
		Type:      models.AffiliationTypeAffiliate.Pointer(),
		ValidFrom: lo.ToPtr(strfmt.DateTime(time.Now())),
	}

	res, err := c.Affiliation.Create(context.Background(), body)

	assert.NoError(t, err)
	assert.NotEmpty(t, res.Data.UID)

	assert.Equal(t, body.OrgUID, res.Data.OrgUID)
	assert.Equal(t, body.PersonUID, res.Data.PersonUID)
	assert.Equal(t, body.Type, res.Data.Type)
	assert.Less(t, time.Time(*body.ValidFrom).Sub(time.Time(*res.Data.ValidFrom)), time.Millisecond)
	assert.Equal(t, body.ValidTo, res.Data.ValidTo)
}

func TestUpdateAffiliation(t *testing.T) {
	c := GetTestClient(ScopePersonAffiliationsWrite, ScopePersonAffiliationsRead)

	currentAff, err := c.Affiliation.Get(context.Background(), TestAffiliationUID)
	assert.NoError(t, err)

	body := models.AffiliationWrite{
		OrgUID:    currentAff.Data.OrgUID,
		PersonUID: currentAff.Data.PersonUID,
		Type:      currentAff.Data.Type,
		ValidFrom: currentAff.Data.ValidFrom,
		ValidTo:   lo.ToPtr(strfmt.DateTime(time.Now())),
	}

	res, err := c.Affiliation.Update(context.Background(), TestAffiliationUID, body)

	assert.NoError(t, err)
	assert.NotEmpty(t, res.Data.UID)

	assert.Equal(t, body.OrgUID, res.Data.OrgUID)
	assert.Equal(t, body.PersonUID, res.Data.PersonUID)
	assert.Equal(t, body.Type, res.Data.Type)
	assert.Less(t, time.Time(*body.ValidFrom).Sub(time.Time(*res.Data.ValidFrom)), time.Millisecond)
	assert.Less(t, time.Time(*body.ValidTo).Sub(time.Time(*res.Data.ValidTo)), time.Millisecond)
}
