package coreapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := getClientForTests(t, ScopePersonsRead)

	token, err := c.tokenSource.Token()
	assert.NoError(t, err)
	assert.NotEmpty(t, token.AccessToken)
}

func getClientForTests(t *testing.T, scopes ...Scope) *Client {
	c := New("http://localhost:3010", WithEmulator(context.Background(), ClientCredentialsEnv{
		TokenUrl: "http://localhost:3020/token",
		Audience: "localhost:3020",
	}, scopes...))

	t.Cleanup(func() { c.RefreshTestData() })

	return c
}
