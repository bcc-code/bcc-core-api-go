package coreapi

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var TestAffiliationUid = uuid.MustParse("8c743b8c-751b-4a76-b48d-92823cd3c44a")

func TestGetAffiliation(t *testing.T) {
	c := GetTestClient(PersonsRead, PersonAffiliationsRead)

	res, err := c.Affiliation.Get(context.Background(), TestAffiliationUid, Fields("*", "person.*"))

	assert.NoError(t, err)
	assert.Equal(t, TestAffiliationUid, res.Data.Uid)
	assert.NotZero(t, TestPersonUid, res.Data.Person)
}

func TestFindAffiliation(t *testing.T) {
	c := GetTestClient(PersonAffiliationsRead)

	res, err := c.Affiliation.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
