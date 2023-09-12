package coreapi

import (
	"net/http"

	"golang.org/x/oauth2"
)

type Client struct {
	http          *http.Client
	apiUrl        string
	apiAudience   string
	oauthTokenUrl string

	tokenSource oauth2.TokenSource
	scope       string

	Person *PersonManager
}

type managerBase struct {
	client *Client
}

func New(url string, options ...ClientOption) *Client {
	c := &Client{
		http:   http.DefaultClient,
		apiUrl: url,
	}

	for _, option := range options {
		option(c)
	}

	if c.tokenSource != nil {
		c.http = addTokenSourceToHttpClient(c.http, c.tokenSource)
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
