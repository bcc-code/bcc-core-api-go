package coreapi

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var TestCountryUid = uuid.MustParse("298aced7-af44-44bf-a21e-c24a110a069f")

func TestGetCountry(t *testing.T) {
	c := GetTestClient(CountriesRead)

	res, err := c.Country.Get(context.Background(), TestCountryUid)

	assert.NoError(t, err)
	assert.Equal(t, TestCountryUid, res.Data.Uid)
}

func TestFindCountry(t *testing.T) {
	c := GetTestClient(CountriesRead)

	res, err := c.Country.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
