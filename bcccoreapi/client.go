package bcccoreapi

import (
	"fmt"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"golang.org/x/oauth2"
)

type Client struct {
	http          *http.Client
	apiUrl        string
	apiAudience   string
	oauthTokenUrl string

	tokenSource oauth2.TokenSource

	httpAgent string

	Affiliation    *genericClient[models.Affiliation, models.AffiliationWrite]
	Consent        *genericClient[models.Consent, models.ConsentWrite]
	Country        *genericClient[models.Country, models.CountryWrite]
	Group          *GroupClient
	Org            *genericClient[models.Org, models.OrgWrite]
	Person         *genericClient[models.Person, models.PersonWrite]
	Relation       *genericClient[models.Relation, models.RelationWrite]
	RoleAssignment *genericClient[models.RoleAssignment, models.RoleAssignmentWrite]
	Role           *genericClient[models.Role, models.RoleWrite]
}

const (
	affiliationsPath    = "/affiliations"
	consentsPath        = "/consents"
	countriesPath       = "/countries"
	groupsPath          = "/groups"
	orgsPath            = "/v2/orgs"
	personsPath         = "/v2/persons"
	relationsPath       = "/relations"
	roleAssignmentsPath = "/roleAssignments"
	rolesPath           = "/roles"
)

var DefaultAgent = fmt.Sprintf("Go-Coreapi/%s", Version)

func NewClient(url string, options ...ClientOption) *Client {
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

	c.Affiliation = &genericClient[models.Affiliation, models.AffiliationWrite]{c, affiliationsPath}
	c.Consent = &genericClient[models.Consent, models.ConsentWrite]{c, consentsPath}
	c.Country = &genericClient[models.Country, models.CountryWrite]{c, countriesPath}
	c.Group = &GroupClient{genericClient[models.Group, models.GroupWrite]{c, groupsPath}}
	c.Org = &genericClient[models.Org, models.OrgWrite]{c, orgsPath}
	c.Person = &genericClient[models.Person, models.PersonWrite]{c, personsPath}
	c.Relation = &genericClient[models.Relation, models.RelationWrite]{c, relationsPath}
	c.RoleAssignment = &genericClient[models.RoleAssignment, models.RoleAssignmentWrite]{c, roleAssignmentsPath}
	c.Role = &genericClient[models.Role, models.RoleWrite]{c, rolesPath}

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
	if w.Base != nil {
		return w.Base.RoundTrip(req)
	}
	return http.DefaultTransport.RoundTrip(req)
}
