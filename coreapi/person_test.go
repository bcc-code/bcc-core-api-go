package coreapi

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestGetPerson(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead)

	res, err := c.Person.Get(context.Background(), TestPersonUID)

	assert.NoError(t, err)
	assert.Equal(t, TestPersonUID, *res.Data.UID)
}

func TestGetPersonNotFound(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead)
	const randomUUID strfmt.UUID = "7ae77f10-c4fb-4b7a-b0c5-ded2c121de4a"

	_, err := c.Person.Get(context.Background(), randomUUID)

	typedErr := &Error{}
	assert.ErrorAs(t, err, &typedErr)
	assert.Equal(t, ErrorCodeNotFound, typedErr.Code)
}

func TestGetPersonResolveFields(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead, ScopeOrgsRead)

	res, err := c.Person.Get(context.Background(), TestPersonUID, Fields("*", "affiliations.*", "affiliations.org.name"))

	assert.NoError(t, err)
	assert.NotEmpty(t, res.Data.Affiliations)
	assert.NotZero(t, res.Data.Affiliations[0].Org.Name)
	assert.Zero(t, res.Data.Affiliations[0].Org.OrgID)
	assert.Zero(t, res.Data.Consents)
}

func TestFindPerson(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead, ScopeOrgsRead)

	res, err := c.Person.Find(context.Background())

	assert.NoError(t, err)
	assert.NotEmpty(t, res.Data)
}

func TestFindPersonFilter(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead, ScopeOrgsRead)

	res, err := c.Person.Find(context.Background(), Filter(fmt.Sprintf(`{"uid": {"_eq": "%s"}}`, TestPersonUID)))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 1)
	assert.Equal(t, TestPersonUID, *res.Data[0].UID)
}

func TestCreateAndUpdatePerson(t *testing.T) {
	c := getClientForTests(t, ScopePersonsWrite)

	body := models.PersonWrite{
		FirstName:   lo.ToPtr("John"),
		LastName:    lo.ToPtr("Doe"),
		DisplayName: lo.ToPtr("John Doe"),
		BirthDate:   lo.ToPtr(strfmt.Date(time.Now().Truncate(24 * time.Hour))),
		Email:       "john@example.com",
		Gender:      models.GenderMale.Pointer(),
		Preferences: &models.Preferences{
			Visibility: &models.VisibilityPreferences{
				Search: models.SearchVisibilityDistrict.Pointer(),
			},
		},
		MaritalStatus: models.MaritalStatusSingle.Pointer(),
	}

	res, err := c.Person.Create(context.Background(), body)
	assert.NoError(t, err)

	t.Run("CheckCreate", func(t *testing.T) {
		assert.NotEmpty(t, res.Data.UID)
		assert.NotZero(t, res.Data.PersonID)

		assert.Equal(t, body.FirstName, res.Data.FirstName)
		assert.Equal(t, body.LastName, res.Data.LastName)
		assert.Equal(t, body.DisplayName, res.Data.DisplayName)
		assert.True(t, body.BirthDate.Equal(*res.Data.BirthDate))
		assert.Equal(t, body.Email, res.Data.Email)
		assert.Equal(t, body.Gender, res.Data.Gender)
		assert.Equal(t, body.Preferences, res.Data.Preferences)
		assert.Equal(t, body.MaritalStatus, res.Data.MaritalStatus)

	})

	t.Run("CheckUpdate", func(t *testing.T) {
		body.Email = "new@example.com"
		res, err := c.Person.Update(context.Background(), *res.Data.UID, body)

		assert.NoError(t, err)
		assert.NotEmpty(t, res.Data.UID)

		assert.Equal(t, body.FirstName, res.Data.FirstName)
		assert.Equal(t, body.LastName, res.Data.LastName)
		assert.Equal(t, body.DisplayName, res.Data.DisplayName)
		assert.True(t, body.BirthDate.Equal(*res.Data.BirthDate))
		assert.Equal(t, body.Email, res.Data.Email)
		assert.Equal(t, body.Gender, res.Data.Gender)
		assert.Equal(t, body.Preferences, res.Data.Preferences)
		assert.Equal(t, body.MaritalStatus, res.Data.MaritalStatus)
	})
}
