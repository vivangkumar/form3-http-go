// Package account exposes account request and response entities.
package account

import (
	"github.com/google/uuid"
)

const accountsType = "accounts"

// Account represents the domain model for a bank account.
type Account struct {
	Attributes     *Attributes `json:"attributes,omitempty"`
	ID             string      `json:"id,omitempty"`
	OrganisationID string      `json:"organisation_id,omitempty"`
	Type           string      `json:"type,omitempty"`
	Version        *int64      `json:"version,omitempty"`
}

// NewAccountWithID returns a builder for Account with a generated
// account ID.
//
// The downstream endpoint will automatically generate UUIDs if one
// is not provided, however this is a convenience method to generate
// an ID client side.
func NewAccountWithID(orgID string) *Account {
	return &Account{
		Type:           accountsType,
		OrganisationID: orgID,
		ID:             uuid.NewString(),
	}
}

// New returns an account with a builder to build the entity.
//
// Only the organisation ID is required.
func New(orgID string) *Account {
	return &Account{
		Type:           accountsType,
		OrganisationID: orgID,
	}
}

// WithID sets the account ID.
func (a *Account) WithID(id string) *Account {
	a.ID = id
	return a
}

// WithOrganisationID sets the organisation ID.
func (a *Account) WithOrganisationID(id string) *Account {
	a.OrganisationID = id
	return a
}

// WithAttributes sets the attributes for an account.
func (a *Account) WithAttributes(attrs *Attributes) *Account {
	a.Attributes = attrs
	return a
}

// Response returns the response from account creation and fetch requests.
type Response struct {
	// Data contains the account returned as part of the response
	Data *Account `json:"data,omitempty"`

	// Links are always returned as part of the response.
	// Except in cases of no content responses.
	Links *Links `json:"links,omitempty"`
}

// DeleteResponse is an empty response type to convey
// a successful delete operation.
type DeleteResponse struct{}

// Links represents the HATEOAS convention links sent as part of responses.
type Links struct {
	Self  string  `json:"self"`
	First *string `json:"first,omitempty"`
	Last  *string `json:"last,omitempty"`
	Next  *string `json:"next,omitempty"`
	Prev  *string `json:"prev,omitempty"`
}
