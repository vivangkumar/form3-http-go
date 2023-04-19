// Package form3 provides the top level interface to interact with the form3 API.
//
// Most consumers of this library should use this package as it presents a unified
// interface to the form3 API.
//
// The client exposes both low level HTTP methods with high level account methods.
package form3

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vivangkumar/form3-http-go/pkg/account"
	accountclient "github.com/vivangkumar/form3-http-go/pkg/account/client"
	baseclient "github.com/vivangkumar/form3-http-go/pkg/client"
)

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

type accountsClient interface {
	Create(
		ctx context.Context,
		acc *account.Account,
	) (*account.Response, error)
	Fetch(
		ctx context.Context,
		params account.FetchAccountParams,
	) (*account.Response, error)
	Delete(
		ctx context.Context,
		params account.DeleteAccountParams,
	) (*account.DeleteResponse, error)
}

// Client represents an abstraction over the base client and the accounts API.
// It exposes a combined interface for callers.
//
// Options may be passed to configure the underlying HTTP client, if required.
type Client struct {
	// Include functionality from the base client.
	baseClient

	// Expose accounts related functionality.
	Accounts accountsClient
}

// New returns a form3 HTTP client.
func New(opts ...Opt) (*Client, error) {
	c, err := baseclient.New(opts...)
	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	return &Client{c, accountclient.New(c)}, nil
}

// Opt aliases client.Opt to delegate application of options
// to the underlying base client.
type Opt = baseclient.Opt
