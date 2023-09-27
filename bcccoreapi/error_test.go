package bcccoreapi

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestCannotParseToken(t *testing.T) {
	c := NewClient(WithCustomEnvironment(EnvironmentConfig{
		BaseUrl: "http://localhost:3010",
	}), WithStaticToken("Invalid"))
	_, err := c.Person.Find(context.Background())
	assertErrType(t, err, ErrorCodeInvalidToken)
}

func TestMissingScopes(t *testing.T) {
	c := getClientForTests(t, ScopeOrgsRead)
	_, err := c.Person.Find(context.Background())
	assertErrType(t, err, ErrorCodeMissingScopes)
}

func TestNotFound(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead)
	const randomUUID strfmt.UUID = "7ae77f10-c4fb-4b7a-b0c5-ded2c121de4a"
	_, err := c.Person.Get(context.Background(), randomUUID)
	assertErrType(t, err, ErrorCodeNotFound)
}

func TestInvalidQuery(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead)
	_, err := c.Person.Find(context.Background(), Fields("invalid"))
	assertErrType(t, err, ErrorCodeInvalidQuery)
}

func assertErrType(t *testing.T, err error, errType ErrorCode) {
	t.Helper()
	typedErr := &Error{}
	assert.ErrorAs(t, err, &typedErr)
	assert.Equal(t, errType, typedErr.Code)
}
