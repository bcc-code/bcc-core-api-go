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

func TestGetConsent(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead, ScopePersonConsentsRead)

	res, err := c.Consent.Get(context.Background(), TestConsentUID, Fields("*", "person.*"))

	assert.NoError(t, err)
	assert.Equal(t, TestConsentUID, *res.Data.UID)
	assert.NotZero(t, TestPersonUID, res.Data.Person)
}

func TestFindConsent(t *testing.T) {
	c := getClientForTests(t, ScopePersonConsentsRead)

	res, err := c.Consent.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}

func TestCreateAndUpdateConsent(t *testing.T) {
	c := getClientForTests(t, ScopePersonConsentsWrite)

	body := models.ConsentWrite{
		OrgUID:    lo.ToPtr(TestOrgUID),
		PersonUID: lo.ToPtr(TestPersonUID),
		Purpose:   models.ConsentPurposeDataSharing.Pointer(),
		ValidFrom: lo.ToPtr(strfmt.DateTime(time.Now().Truncate(time.Millisecond))),
	}

	res, err := c.Consent.Create(context.Background(), body)
	assert.NoError(t, err)

	t.Run("CheckCreate", func(t *testing.T) {
		assert.NotEmpty(t, res.Data.UID)

		assert.Equal(t, body.OrgUID, res.Data.OrgUID)
		assert.Equal(t, body.PersonUID, res.Data.PersonUID)
		assert.Equal(t, body.Purpose, res.Data.Purpose)
		assert.True(t, body.ValidFrom.Equal(*res.Data.ValidFrom))
		assert.Equal(t, body.ValidTo, res.Data.ValidTo)

	})

	t.Run("CheckUpdate", func(t *testing.T) {
		body.ValidTo = lo.ToPtr(strfmt.DateTime(time.Now().Truncate(time.Millisecond)))
		res, err := c.Consent.Update(context.Background(), *res.Data.UID, body)

		assert.NoError(t, err)
		assert.NotEmpty(t, res.Data.UID)

		assert.Equal(t, body.OrgUID, res.Data.OrgUID)
		assert.Equal(t, body.PersonUID, res.Data.PersonUID)
		assert.Equal(t, body.Purpose, res.Data.Purpose)
		assert.True(t, body.ValidFrom.Equal(*res.Data.ValidFrom))
		assert.True(t, body.ValidTo.Equal(*res.Data.ValidTo))

	})
}
