package coreapi

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

type Client struct {
	http          *http.Client
	apiUrl        string
	apiAudience   string
	oauthTokenUrl string

	tokenSource oauth2.TokenSource

	httpAgent string

	Person *PersonManager
}

type managerBase struct {
	client *Client
}

var DefaultAgent = fmt.Sprintf("Go-Coreapi/%s", Version)

func New(url string, options ...ClientOption) *Client {
	c := &Client{
		http:      http.DefaultClient,
		apiUrl:    url,
		httpAgent: DefaultAgent,
	}

	for _, option := range options {
		option(c)
	}

	if c.tokenSource != nil {
		c.http = addTokenSourceToHttpClient(c.http, c.tokenSource)
	}

	c.http.Transport = wrappedRoundTripper{
		Base: c.http.Transport,
		Fn: func(req *http.Request) error {
			req.Header.Set("User-Agent", c.httpAgent)
			return nil
		},
	}

	base := &managerBase{c}

	c.Person = (*PersonManager)(base)

	return c
}

func addTokenSourceToHttpClient(base *http.Client, ts oauth2.TokenSource) *http.Client {
	return &http.Client{
		Timeout: base.Timeout,
		Transport: &oauth2.Transport{
			Base:   base.Transport,
			Source: ts,
		},
	}
}

type wrappedRoundTripper struct {
	Base http.RoundTripper
	Fn   func(*http.Request) error
}

func (w wrappedRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	err := w.Fn(req)
	if err != nil {
		return nil, err
	}
	return w.Base.RoundTrip(req)
}
