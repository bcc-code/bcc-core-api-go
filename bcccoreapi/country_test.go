package bcccoreapi

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

const TestCountryUID strfmt.UUID = "298aced7-af44-44bf-a21e-c24a110a069f"

func TestGetCountry(t *testing.T) {
	c := getClientForTests(t, ScopeCountriesRead)

	res, err := c.Country.Get(context.Background(), TestCountryUID)

	assert.NoError(t, err)
	assert.Equal(t, TestCountryUID, *res.Data.UID)
}

func TestFindCountry(t *testing.T) {
	c := getClientForTests(t, ScopeCountriesRead)

	res, err := c.Country.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
