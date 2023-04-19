# form3-http-go

Go HTTP client library to access the form3 accounts API.

**This was built as part of the form3 take home exercise as part of the interview process**

## Limitations

Currently, functionality is limited only to `Create`, `Fetch` and `Delete` accounts as mentioned in the [submission guidance](https://github.com/form3tech-oss/interview-accountapi#submission-guidance).

## Usage

```
go get github.com/vivangkumar/form3-http-go
```

Extremely basic usage of the library 👇

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vivangkumar/form3-http-go/pkg/account"
	"github.com/vivangkumar/form3-http-go/pkg/form3"
)

func main() {
	// Simplest initialisation of the client.
	client, err := form3.New()
	if err != nil {
		log.Fatalf(err.Error())
	}

	ctx := context.Background()

	attrs := account.
		NewAttributes("currency", "country").
		WithBankID("bank-id").
		WithBankIDCode("bank-id-code")

	acc := account.New("org-id").WithAttributes(attrs)

	resp, err := client.Accounts.Create(ctx, acc)
	if err != nil {
		log.Fatalf("create account: %s", err.Error())
	}

	fmt.Println("created account with ID: ", resp.Data.ID)
}
```

## Configuration

The client can be configured with various options to customise its behaviour.

### WithHTTPClient

This can be used to specify a custom HTTP client.

For example: to set your own transport and timeout on the client 👇

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/vivangkumar/form3-http-go/pkg/client"
	"github.com/vivangkumar/form3-http-go/pkg/form3"
)

func main() {
	_, err := form3.New(
		client.WithHTTPClient(
			&http.Client{
				Transport: &http.Transport{},
				Timeout:   1 * time.Second,
			},
		),
	)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
```

### WithBaseURL

This can be used to specify a different base URL to the client.
By default, requests are sent to `https://api.form3.tech`

This is useful to develop locally.

```go
package main

import (
	"log"

	"github.com/vivangkumar/form3-http-go/pkg/client"
	"github.com/vivangkumar/form3-http-go/pkg/form3"
)

func main() {
	_, err := form3.New(
		client.WithBaseURL("http://localhost:8080"),
	)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
```

### WithHTTPRequestHeaders

This can be used to set custom headers on HTTP requests from the client.

For example: to change the user agent 👇

```go
package main

import (
	"log"

	"github.com/vivangkumar/form3-http-go/pkg/client"
	"github.com/vivangkumar/form3-http-go/pkg/form3"
)

func main() {
	_, err := form3.New(
		client.WithHTTPRequestHeaders(map[string]string{
			"User-Agent": "mylibrary/v1",
		}),
	)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
```

## Base client

The base client acts as the entry point to make requests to the form3 API.

The base client exposes two distinct types of request making behaviour.
1. Raw requests to the API via `NewRequest`, `Do` and `Get`, `Post`, `Delete`.
2. Accessing specific resources directly (via `Accounts`)

## Accounts API

Usage example:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vivangkumar/form3-http-go/pkg/account"
	"github.com/vivangkumar/form3-http-go/pkg/form3"
)

func main() {
	client, err := form3.New()
	if err != nil {
		log.Fatalf(err.Error())
	}

	ctx := context.Background()

	attrs := account.
		NewAttributes("currency", "country").
		WithBankID("bank-id").
		WithBankIDCode("bank-id-code")

	acc := account.New("org-id").WithAttributes(attrs)

	// Create account.
	created, err := client.Accounts.Create(ctx, acc)
	if err != nil {
		log.Fatalf("create account: %s", err.Error())
	}

	fmt.Println(created.Data)
	fmt.Println(created.Links)

	// Fetch account.
	_, err = client.Accounts.Fetch(ctx, account.FetchAccountParams{
		ID: created.ID,
	})
	if err != nil {
		log.Fatalf("fetch account: %s", err.Error())
	}

	// Delete account.
	_, err = client.Accounts.Delete(ctx, account.DeleteAccountParams{
		ID:      created.ID,
		Version: 1,
	})
	if err != nil {
		log.Fatalf("delete account: %s", err.Error())
	}
}
```

### Inspecting API errors

In the simplest form, the returned error should carry enough details sufficient for logging and adding context to other callers up the stack.

However, to peek into the HTTP response, it is possible to inspect the returned error.
The library does not return concrete error types and instead requires errors to be
treated as opaque.

However, errors returned will satisfy the below interface:

```go
type errResponse interface {
	HTTPResponse() *http.Response
	Error() string
}
```

Clients can get more details by using `errors.As()` to get more details.

```go
package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/vivangkumar/form3-http-go/pkg/account"
	"github.com/vivangkumar/form3-http-go/pkg/form3"
)

func main() {
	client, err := form3.New()
	if err != nil {
		log.Fatalf(err.Error())
	}

	ctx := context.Background()

	attrs := account.
		NewAttributes("currency", "country").
		WithBankID("bank-id").
		WithBankIDCode("bank-id-code")

	acc := account.New("org-id").WithAttributes(attrs)

	// Create account.
	_, err = client.Accounts.Create(ctx, acc)
	if err != nil {
		var e errResponse
		if errors.As(err, &e) {
			// Get the underlying HTTP response.
			resp := e.HTTPResponse()
			log.Fatalf("create account: %s (%d)", e.Error(), resp.StatusCode)
		}
	}
}

type errResponse interface {
	HTTPResponse() *http.Response
	Error() string
}
```

## Tests

The code is covered by both unit and integration tests.

These tests also run as part of `docker-compose up`.

To run all the tests, use

```
make test
```

To run just the unit tests, use

```
make unit-test
```

Finally, to run integration tests,

```
ACCOUNTS_API_BASE_URL=http://localhost:8080 make integration-test
```

This will run against a fake account API running at `ACCOUNTS_API_BASE_URL`

## Package structure

A flexible package structure has been used where each package can be used in isolation if required.

The `form3` package brings together functionality from the other two packages
and as such acts as a wrapper over the other two.

- `account` includes all entities required to interact with the accounts endpoints.
	It also provides an account client that can be used to interact solely
	with the accounts API.
- `client` presents a low-level HTTP client that is used by the account client.
	This client can also be used to make requests to the API without relying on
  response types being returned.
- `form3` presents a unified interface to the above two packages.
  Most callers should use this package.
