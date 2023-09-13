package coreapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := New("http://localhost:3010", WithEmulator(context.Background(), ClientCredentialsEnv{
		TokenUrl: "http://localhost:3020/token",
		Audience: "localhost:3020",
	}, "persons#read"))
	p, err := c.Person.Find(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, p.Data)
}
