package client_test

import (
	"context"
	"io"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vivangkumar/form3-http-go/pkg/client"
	"github.com/vivangkumar/form3-http-go/pkg/client/internal/fakes"
)

var _ = Describe("Setting client options", func() {
	var (
		fakeHTTPClient *fakes.FakeHttpClient
		cl             *client.Client

		ctx    context.Context
		cancel context.CancelFunc

		baseURL string
	)

	BeforeEach(func() {
		fakeHTTPClient = new(fakes.FakeHttpClient)
		fakeHTTPClient.DoReturns(&http.Response{StatusCode: http.StatusOK}, nil)

		ctx, cancel = context.WithCancel(context.Background())

		baseURL = "http://localhost:8080"
	})

	When("a client is configured", func() {
		Context("with valid configuration", func() {
			BeforeEach(func() {
				c, err := client.New(
					client.WithHTTPClient(fakeHTTPClient),
					client.WithBaseURL(baseURL),
					client.WithHTTPRequestHeaders(map[string]string{
						"User-Agent": "client/go",
					}),
				)
				Expect(err).To(BeNil())
				cl = c

				fakeHTTPClient.DoReturns(&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(http.NoBody),
				}, nil)
			})

			It("should make HTTP requests", func() {
				path := "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

				req, err := cl.NewRequest(ctx, http.MethodGet, path, nil, nil)
				Expect(err).To(BeNil())

				resp, err := cl.Do(req, nil)
				Expect(err).To(BeNil())
				defer cancel()

				By("using the configured http client", func() {
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
					Expect(fakeHTTPClient.DoCallCount()).To(Equal(1))
				})

				By("using the configured base URL", func() {
					Expect(req.Host).To(Equal("localhost:8080"))
				})

				By("using the configured headers as part of request headers", func() {
					Expect(req.Header.Get("User-Agent")).To(Equal("client/go"))
				})
			})
		})

		Context("with invalid base URL", func() {
			var err error

			BeforeEach(func() {
				baseURL = "http:localhost"

				_, err = client.New(
					client.WithHTTPClient(fakeHTTPClient),
					client.WithBaseURL(baseURL),
					client.WithHTTPRequestHeaders(map[string]string{
						"User-Agent": "client/go",
					}),
				)
			})

			It("should return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
