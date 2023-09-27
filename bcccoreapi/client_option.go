package bcccoreapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/samber/lo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type ClientOption func(*Client)

func WithEnvironment(env Environment) ClientOption {
	envConfig, ok := envMap[env]
	if !ok {
		panic("invalid environment")
	}

	return WithCustomEnvironment(envConfig)
}

func WithCustomEnvironment(envConfig EnvironmentConfig) ClientOption {
	return func(c *Client) {
		c.envConfig = envConfig
	}
}

// WithClientCredentials configures the SDK to authenticate using the client
// credentials authentication flow.
func WithClientCredentials(ctx context.Context, clientID, clientSecret string, scopes ...Scope) ClientOption {
	return func(c *Client) {
		cfg := &clientcredentials.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			TokenURL:     c.envConfig.TokenUrl,
			Scopes:       scopesToStrings(scopes),
			EndpointParams: url.Values{
				"audience": []string{c.envConfig.Audience},
			},
		}

		c.tokenSource = cfg.TokenSource(ctx)
	}
}

func scopesToStrings(scopes []Scope) []string {
	return lo.Map(scopes, func(s Scope, _ int) string {
		return string(s)
	})
}

// WithClientCredentials configures the SDK to authenticate using the client
// credentials authentication flow against the emulator.
func WithEmulator(ctx context.Context, scopes ...Scope) ClientOption {
	return func(c *Client) {
		cfg := &clientcredentials.Config{
			ClientID:     "",
			ClientSecret: "",
			TokenURL:     c.envConfig.TokenUrl,
			Scopes:       scopesToStrings(scopes),
			EndpointParams: url.Values{
				"audience":      []string{c.envConfig.Audience},
				"custom_claims": []string{`{"orgs": "-", "app_uid": "990dc470-c4f4-429e-a11e-f094cb1cf7c0"}`},
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

// WithTokenSource configures the SDK to authenticate using the provided token source
func WithTokenSource(tokenSource oauth2.TokenSource) ClientOption {
	return func(c *Client) {
		c.tokenSource = tokenSource
	}
}

// WithHTTPClient configures the SDK to use the provided client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(m *Client) {
		m.http = client
	}
}
