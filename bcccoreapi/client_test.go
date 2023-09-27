package bcccoreapi

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
	c := NewClient(WithCustomEnvironment(EnvironmentConfig{
		BaseUrl:  "http://localhost:3010",
		TokenUrl: "http://localhost:3020/token",
		Audience: "localhost:3020",
	}), WithEmulator(context.Background(), scopes...))

	t.Cleanup(func() { c.RefreshTestData() })

	return c
}
