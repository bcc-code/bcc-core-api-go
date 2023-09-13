//go:generate go-enum --marshal

package coreapi

import (
	"context"
	"net/http"
	"time"

	"cloud.google.com/go/civil"
	"github.com/google/uuid"
)

// replace (json:"[a-zA-Z0-9]*").*

type Person struct {
	Uid             uuid.UUID        `json:"uid"`
	PersonID        int              `json:"personID"`
	Affiliations    []Affiliation    `json:"affiliations"`
	Relations       []PersonRelation `json:"relations"`
	Consents        []Consent        `json:"consents"`
	RoleAssignments []RoleAssignment `json:"roleAssignments"`
	LastChangedDate time.Time        `json:"lastChangedDate"`

	FirstName         string        `json:"firstName"`
	MiddleName        string        `json:"middleName"`
	LastName          string        `json:"lastName"`
	DisplayName       string        `json:"displayName"`
	BirthDate         civil.Date    `json:"birthDate"`
	DeceasedDate      civil.Date    `json:"deceasedDate"`
	Gender            Gender        `json:"gender"`
	MaritalStatus     MaritalStatus `json:"maritalStatus"`
	Email             string        `json:"email"`
	EmailVerified     bool          `json:"emailVerified"`
	CellPhone         string        `json:"cellPhone"`
	CellPhoneVerified bool          `json:"cellPhoneVerified"`
	HomePhone         string        `json:"homePhone"`
	Preferences       Preferences   `json:"preferences"`
	Address           Address       `json:"address"`
	NationalIds       []NationalId  `json:"nationalIds"`

	// URL of person's profile picture
	ProfilePicture string `json:"profilePicture"`
}

// ENUM(Male, Female, Unknown)
type Gender string

// ENUM(Single, Married, Widowed, Separated, SingleParent, Unknown)
type MaritalStatus string

type Address struct {
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	Address3    string `json:"address3"`
	Region      string `json:"region"`
	City        string `json:"city"`
	PostalCode  string `json:"postalCode"`
	CountryCode string `json:"countryCode"`
}

type Preferences struct {
	ContentLanguages []string              `json:"contentLanguages"`
	UiLanguages      []string              `json:"uiLanguages"`
	Visibility       VisibilityPreferences `json:"visibility"`
}

type VisibilityPreferences struct {
	BirthdayList bool             `json:"birthdayList"`
	Search       SearchVisibility `json:"search"`
}

// ENUM(Global, District, Hidden)
type SearchVisibility string

type NationalId struct {
	CountryIso2Code string `json:"countryIso2Code"`
	Id              string `json:"id"`
}

type PersonRelation struct {
	TargetUid uuid.UUID `json:"targetUid"`
	Target    *Person   `json:"target"`

	// Permission of origin on target
	GrantToTarget Grant `json:"grantToTarget"`
	// Permission of target on origin
	GrantToOrigin Grant `json:"grantToOrigin"`

	// Type of relation, defined as {target} is {type} of {origin}
	Type      PersonRelationType `json:"type"`
	ValidFrom time.Time          `json:"validFrom"`
	ValidTo   *time.Time         `json:"validTo"`
} //@name PersonRelation

// ENUM(Child, Parent, Spouse, LegalDependent, LegalGuardian, FosterChild, FosterParent, ContactDependent, ContactPerson)
type PersonRelationType string //@name PersonRelationType

var PersonPath = "/v2/persons"

type PersonManager managerBase

func (m *PersonManager) Get(ctx context.Context, uid uuid.UUID, opts ...RequestOption) (Response[Person], error) {
	var res Response[Person]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(PersonPath, uid.String()), nil, &res, opts...)
	return res, err
}

func (m *PersonManager) Find(ctx context.Context, opts ...RequestOption) (ResponseWithMeta[[]Person], error) {
	var res ResponseWithMeta[[]Person]
	err := m.client.Request(ctx, http.MethodGet, m.client.URL(PersonPath), nil, &res, opts...)
	return res, err
}
