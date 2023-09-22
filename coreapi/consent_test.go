package coreapi

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

const TestConsentUID strfmt.UUID = "6026dbaa-ed8e-4fc8-993f-8ce377b4cfee"

func TestGetConsent(t *testing.T) {
	c := GetTestClient(ScopePersonsRead, ScopePersonConsentsRead)

	res, err := c.Consent.Get(context.Background(), TestConsentUID, Fields("*", "person.*"))

	assert.NoError(t, err)
	assert.Equal(t, TestConsentUID, *res.Data.UID)
	assert.NotZero(t, TestPersonUID, res.Data.Person)
}

func TestFindConsent(t *testing.T) {
	c := GetTestClient(ScopePersonConsentsRead)

	res, err := c.Consent.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
