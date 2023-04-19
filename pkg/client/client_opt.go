package client

import (
	"fmt"
	"net/url"
)

// Opt represents an option that can be passed
// during creation of a Client to configure it.
type Opt func(c *Client) error

// WithHTTPClient sets the underlying HTTP client.
//
// It must conform to the httpClient interface.
//
// Use this for advanced configuration of the HTTP
// client.
func WithHTTPClient(hc httpClient) Opt {
	return func(c *Client) error {
		c.httpClient = hc
		return nil
	}
}

// WithBaseURL sets the base URL for the client.
//
// If not used, the default base URL is used.
func WithBaseURL(base string) Opt {
	return func(c *Client) error {
		u, err := url.Parse(base)
		if err != nil {
			return fmt.Errorf("base url opt: %w", err)
		}
		c.baseURL = u

		return nil
	}
}

// WithHTTPRequestHeaders sets optional HTTP headers for each request.
func WithHTTPRequestHeaders(headers map[string]string) Opt {
	return func(c *Client) error {
		for k, v := range headers {
			c.headers[k] = v
		}

		return nil
	}
}
