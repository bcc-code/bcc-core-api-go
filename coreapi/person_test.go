package coreapi

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestGetPerson(t *testing.T) {
	c := GetTestClient(ScopePersonsRead)

	res, err := c.Person.Get(context.Background(), TestPersonUID)

	assert.NoError(t, err)
	assert.Equal(t, TestPersonUID, *res.Data.UID)
}

func TestGetPersonNotFound(t *testing.T) {
	c := GetTestClient(ScopePersonsRead)
	const randomUUID strfmt.UUID = "7ae77f10-c4fb-4b7a-b0c5-ded2c121de4a"

	_, err := c.Person.Get(context.Background(), randomUUID)

	typedErr := &Error{}
	assert.ErrorAs(t, err, &typedErr)
	assert.Equal(t, ErrorCodeNotFound, typedErr.Code)
}

func TestGetPersonResolveFields(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopeOrgsRead)

	res, err := c.Person.Get(context.Background(), TestPersonUID, Fields("*", "affiliations.*", "affiliations.org.name"))

	assert.NoError(t, err)
	assert.NotEmpty(t, res.Data.Affiliations)
	assert.NotZero(t, res.Data.Affiliations[0].Org.Name)
	assert.Zero(t, res.Data.Affiliations[0].Org.OrgID)
	assert.Zero(t, res.Data.Consents)
}

func TestFindPerson(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopeOrgsRead)

	res, err := c.Person.Find(context.Background())

	assert.NoError(t, err)
	assert.NotEmpty(t, res.Data)
}

func TestFindPersonFilter(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopeOrgsRead)

	res, err := c.Person.Find(context.Background(), Filter(fmt.Sprintf(`{"uid": {"_eq": "%s"}}`, TestPersonUID)))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 1)
	assert.Equal(t, TestPersonUID, *res.Data[0].UID)
}
