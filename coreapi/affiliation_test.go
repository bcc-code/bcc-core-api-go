package coreapi

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

const TestAffiliationUID strfmt.UUID = "8c743b8c-751b-4a76-b48d-92823cd3c44a"

func TestGetAffiliation(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopePersonAffiliationsRead)

	res, err := c.Affiliation.Get(context.Background(), TestAffiliationUID, Fields("*", "person.*"))

	assert.NoError(t, err)
	assert.Equal(t, TestAffiliationUID, *res.Data.UID)
	assert.NotZero(t, TestPersonUID, res.Data.Person)
}

func TestFindAffiliation(t *testing.T) {
	c := GetTestClient(ScopePersonAffiliationsRead)

	res, err := c.Affiliation.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
