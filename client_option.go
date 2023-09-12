package coreapi

import (
	"context"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type ClientOption func(*Client)

// WithClientCredentials configures the SDK to authenticate using the client
// credentials authentication flow.
func WithClientCredentials(ctx context.Context, credEnv ClientCredentialsEnv, clientID, clientSecret string, scopes ...string) ClientOption {
	return func(c *Client) {
		cfg := &clientcredentials.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			TokenURL:     credEnv.TokenUrl,
			Scopes:       scopes,
			EndpointParams: url.Values{
				"audience": []string{credEnv.Audience},
			},
		}

		c.tokenSource = cfg.TokenSource(ctx)
	}
}

// WithClientCredentials configures the SDK to authenticate using the client
// credentials authentication flow.
func WithEmulator(ctx context.Context, credEnv ClientCredentialsEnv, scopes ...string) ClientOption {
	return func(c *Client) {
		cfg := &clientcredentials.Config{
			ClientID:     "",
			ClientSecret: "",
			TokenURL:     credEnv.TokenUrl,
			Scopes:       scopes,
			EndpointParams: url.Values{
				"audience":      []string{credEnv.Audience},
				"custom_claims": []string{`{"orgs": "-"}`},
			},
		}

		c.tokenSource = cfg.TokenSource(ctx)
	}
}

// WithClientCredentials configures the SDK to authenticate using the provided token
func WithStaticToken(token string) ClientOption {
	return func(c *Client) {
		c.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	}
}

// WithClientCredentials configures the SDK to authenticate using the provided token source
func WithTokenSource(tokenSource oauth2.TokenSource) ClientOption {
	return func(c *Client) {
		c.tokenSource = tokenSource
	}
}

// WithClient configures the SDK to use the provided client.
func WithClient(client *http.Client) ClientOption {
	return func(m *Client) {
		m.http = client
	}
}
