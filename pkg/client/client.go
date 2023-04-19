// Package client provides functionality to make HTTP requests to the form3 API.
//
// It provides convenience methods to make HTTP requests and handles creating requests
// with the correct endpoints and setting appropriate headers.
//
// Additionally, it provides configuration options for the client that can be used to
// configure various aspects.
//
// This package can be used on its own or can be used to configure the accounts client.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultTimeout = 3 * time.Second
	defaultBaseURL = "https://api.form3.tech"

	clientVersion = "0.1.0"
	userAgent     = "form3-http-go/" + clientVersion

	contentTypeHeader = "application/vnd.api+json"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o internal/fakes/fake_http_client.go . httpClient
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client represents a form3 HTTP API client.
type Client struct {
	// httpClient is the underlying http client.
	httpClient httpClient

	// baseURL is the base form3 http endpoint.
	//
	// Requests to other paths are constructed from the base path.
	baseURL *url.URL

	// headers represent optional headers to set on each request.
	headers map[string]string
}

// New constructs a form3 http client.
//
// It uses the default base path when constructed.
//
// Additionally, options may be passed to configure the client.
func New(opts ...Opt) (*Client, error) {
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		httpClient: &http.Client{
			Timeout:   defaultTimeout,
			Transport: &http.Transport{},
		},
		baseURL: baseURL,
		headers: make(map[string]string),
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, fmt.Errorf("apply options: %w", err)
		}
	}

	return c, nil
}

// Get is a convenience to create and execute a GET request against the API.
//
// Result from the API is decoded into target.
func (c *Client) Get(
	ctx context.Context,
	path string,
	query map[string]string,
	target any,
) (*http.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, path, query, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	return c.Do(req, target)
}

// Post is a convenience method to create and execute a POST request against the API.
//
// body represents the request body, while target is the value to which an API
// response will be decoded into.
func (c *Client) Post(
	ctx context.Context,
	path string,
	body any,
	target any,
) (*http.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, path, nil, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	return c.Do(req, target)
}

// Delete is a convenience method to create and execute a DELETE request against the API.
func (c *Client) Delete(
	ctx context.Context,
	path string,
	query map[string]string,
) (*http.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodDelete, path, query, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	return c.Do(req, nil)
}

// NewRequest creates an HTTP request, but additionally also
// handles constructing the endpoint and setting required headers.
func (c *Client) NewRequest(
	ctx context.Context,
	method string,
	path string,
	query map[string]string,
	body any,
) (*http.Request, error) {
	// Add a prefixed / if one doesn't exist.
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	u, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}

	// Set query parameters.
	if len(query) > 0 {
		q := u.Query()
		for k, v := range query {
			q.Add(k, v)
		}

		u.RawQuery = q.Encode()
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, fmt.Errorf("encode body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// Set content type header.
	if body != nil {
		req.Header.Set("Content-Type", contentTypeHeader)
	}

	// Set required headers for all requests.
	req.Header.Set("Host", c.baseURL.Host)
	req.Header.Set("Date", time.Now().UTC().Format(time.RFC850))
	req.Header.Set("Accept", contentTypeHeader)
	req.Header.Set("User-Agent", userAgent)

	// Add additional headers, if any.
	if len(c.headers) > 0 {
		for k, v := range c.headers {
			req.Header.Set(k, v)
		}
	}

	return req, nil
}

// Do makes an HTTP request and returns the response from the API.
//
// The returned response is JSON decoded into the value pointed to by target.
//
// If an API error has occurred i.e where the status code is not 2xx, then an
// ErrorResponse is returned.
func (c *Client) Do(req *http.Request, target any) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	err = c.maybeDecodeAPIError(resp)
	if err != nil {
		return nil, err
	}

	if target != nil && resp.StatusCode != http.StatusNoContent {
		err := json.NewDecoder(resp.Body).Decode(target)
		if err != nil {
			return nil, fmt.Errorf("decode response body: %w", err)
		}
	}

	return resp, nil
}

// maybeDecodeAPIError checks for any errors returned as part of the response body.
//
// If there is one, the body is JSON decoded and an error message is constructed.
func (c *Client) maybeDecodeAPIError(resp *http.Response) error {
	sc := resp.StatusCode
	// These are the only success codes reported by the API.
	if sc >= 200 && sc <= 204 {
		return nil
	}

	r := errorResponse{httpResponse: resp}
	// 400, 409 and 403 return JSON response bodies
	// https://www.api-docs.form3.tech/api/schemes/sepa-instant-credit-transfer/introduction/errors-status-codes
	if sc == http.StatusBadRequest || sc == http.StatusConflict ||
		sc == http.StatusForbidden {
		var e apiError
		err := json.NewDecoder(resp.Body).Decode(&e)
		if err != nil {
			return fmt.Errorf(
				"decode error response (HTTP status: %d): %w",
				resp.StatusCode,
				err,
			)
		}

		r.underlying = &e
	}

	return &r
}
