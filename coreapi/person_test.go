package coreapi

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var TestPersonUid = uuid.MustParse("657a66ca-9cd0-4b61-9476-697016e26fbc")

func TestGetPerson(t *testing.T) {
	c := GetTestClient(ScopePersonsRead)

	res, err := c.Person.Get(context.Background(), TestPersonUid)

	assert.NoError(t, err)
	assert.Equal(t, TestPersonUid, res.Data.Uid)
}

func TestGetPersonNotFound(t *testing.T) {
	c := GetTestClient(ScopePersonsRead)

	_, err := c.Person.Get(context.Background(), uuid.New())

	typedErr := &Error{}
	assert.ErrorAs(t, err, &typedErr)
	assert.Equal(t, ErrorCodeNotFound, typedErr.Code)
}

func TestGetPersonResolveFields(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopeOrgsRead)

	res, err := c.Person.Get(context.Background(), TestPersonUid, Fields("*", "affiliations.*", "affiliations.org.name"))

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

	res, err := c.Person.Find(context.Background(), Filter(fmt.Sprintf(`{"uid": {"_eq": "%s"}}`, TestPersonUid)))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 1)
	assert.Equal(t, TestPersonUid, res.Data[0].Uid)
}
