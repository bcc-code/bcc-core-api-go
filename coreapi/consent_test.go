package coreapi

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var TestConsentUid = uuid.MustParse("6026dbaa-ed8e-4fc8-993f-8ce377b4cfee")

func TestGetConsent(t *testing.T) {
	c := GetTestClient(PersonsRead, PersonConsentsRead)

	res, err := c.Consent.Get(context.Background(), TestConsentUid, Fields("*", "person.*"))

	assert.NoError(t, err)
	assert.Equal(t, TestConsentUid, res.Data.Uid)
	assert.NotZero(t, TestPersonUid, res.Data.Person)
}

func TestFindConsent(t *testing.T) {
	c := GetTestClient(PersonConsentsRead)

	res, err := c.Consent.Find(context.Background(), Limit(2))

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}
