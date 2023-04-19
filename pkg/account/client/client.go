// Package client provides functionality to interact with the form3 accounts API.
//
// To use the client in this package, a baseClient is required. The client exported
// from the client package is suitable for use here.
//
// Requests made via this client are made against the /v1/organisations/accounts/
// endpoints.

package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vivangkumar/form3-http-go/pkg/account"
)

const accountsBasePath = "/v1/organisation/accounts/"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o internal/fakes/fake_base_client.go . baseClient
type baseClient interface {
	Get(
		ctx context.Context,
		path string,
		query map[string]string,
		target any,
	) (*http.Response, error)
	Post(
		ctx context.Context,
		path string,
		body any,
		target any,
	) (*http.Response, error)
	Delete(
		ctx context.Context,
		path string,
		query map[string]string,
	) (*http.Response, error)
}

type accountCreationRequest struct {
	Data *account.Account `json:"data,omitempty"`
}

// Client represents an account client.
type Client struct {
	baseClient baseClient
}

// New creates a new account client.
//
// It requires an underlying client that satisfies the baseClient interface.
func New(baseClient baseClient) *Client {
	return &Client{baseClient: baseClient}
}

// Create creates a new bank account.
func (c *Client) Create(
	ctx context.Context,
	acc *account.Account,
) (*account.Response, error) {
	if acc == nil {
		return nil, fmt.Errorf("account entity is nil")
	}

	target := new(account.Response)
	req := accountCreationRequest{Data: acc}

	_, err := c.baseClient.Post(ctx, accountsBasePath, &req, target)
	if err != nil {
		return nil, fmt.Errorf("create account: %w", err)
	}

	return target, nil
}

// Fetch retrieves an account from the API given an ID.
func (c *Client) Fetch(
	ctx context.Context,
	params account.FetchAccountParams,
) (*account.Response, error) {
	target := new(account.Response)

	_, err := c.baseClient.Get(ctx, accountsBasePath+params.ID, nil, target)
	if err != nil {
		return nil, fmt.Errorf("fetch account: %w", err)
	}

	return target, nil
}

// Delete deletes the account with the given ID and version.
func (c *Client) Delete(
	ctx context.Context,
	params account.DeleteAccountParams,
) (*account.DeleteResponse, error) {
	query := map[string]string{
		"version": fmt.Sprintf("%d", params.Version),
	}

	_, err := c.baseClient.Delete(ctx, accountsBasePath+params.ID, query)
	if err != nil {
		return nil, fmt.Errorf("delete account: %w", err)
	}

	return &account.DeleteResponse{}, nil
}
