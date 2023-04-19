package client_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vivangkumar/form3-http-go/pkg/account"
	"github.com/vivangkumar/form3-http-go/pkg/client"
	"github.com/vivangkumar/form3-http-go/pkg/client/internal/fakes"
	"github.com/vivangkumar/form3-http-go/pkg/internal/fixtures"
)

var _ = Describe("Client", func() {
	var (
		cl             *client.Client
		fakeHTTPClient *fakes.FakeHttpClient

		host string

		basePath string
		path     string

		orgID     string
		accountID string

		ctx context.Context
	)

	BeforeEach(func() {
		ctx = context.Background()

		fakeHTTPClient = new(fakes.FakeHttpClient)

		cl, _ = client.New(client.WithHTTPClient(fakeHTTPClient))

		orgID = uuid.NewString()
		accountID = uuid.NewString()

		host = "api.form3.tech"
		basePath = "/v1/organisation"
		path = basePath
	})

	Describe("General request construction", func() {
		BeforeEach(func() {
			path = path + "/accounts/" + accountID
		})

		Context("with an invalid request path", func() {
			BeforeEach(func() {
				path = "bad\npath"
			})

			It("should return an error", func() {
				req, err := cl.NewRequest(ctx, http.MethodGet, path, nil, nil)
				Expect(err).To(Not(BeNil()))
				Expect(req).To(BeNil())
			})
		})

		Context("with additional headers configured on the client", func() {
			BeforeEach(func() {
				err := client.WithHTTPRequestHeaders(map[string]string{
					"My-Header": "Hello",
				})(cl)
				Expect(err).To(BeNil())
			})

			It("should add the headers as part of the request", func() {
				req, err := cl.NewRequest(ctx, http.MethodGet, path, nil, nil)
				Expect(err).To(BeNil())

				Expect(req.Header.Get("My-Header")).To(Equal("Hello"))
			})
		})

		Context("with query parameters", func() {
			It("should encode query parameters in the request", func() {
				params := map[string]string{"version": "1"}

				req, err := cl.NewRequest(ctx, http.MethodGet, path, params, nil)
				Expect(err).To(BeNil())
				
				Expect(req.URL.Query().Get("version")).To(Equal("1"))
				Expect(req.URL.String()).To(Equal("https://" + host + path + "?version=1"))
			})
		})
	})

	Describe("Creating a GET request", func() {
		BeforeEach(func() {
			path = path + "/accounts/" + accountID
		})

		It("should return a GET request", func() {
			req, err := cl.NewRequest(ctx, http.MethodGet, path, nil, nil)
			Expect(err).To(BeNil())

			By("setting the correct request URL", func() {
				Expect(req.URL.String()).To(Equal("https://" + host + path))
			})

			By("setting the correct HTTP method", func() {
				Expect(req.Method).To(Equal(http.MethodGet))
			})

			By("not including a request body", func() {
				Expect(req.Body).To(BeNil())
			})

			By("setting the default headers", func() {
				Expect(req.Header.Get("Host")).To(Equal(host))
				Expect(req.Header.Get("Accept")).To(Equal("application/vnd.api+json"))
				Expect(req.Header.Get("Date")).To(Not(BeEmpty()))
				Expect(req.Header.Get("User-Agent")).To(Not(BeEmpty()))
			})
		})
	})

	Describe("Creating a POST request", func() {
		BeforeEach(func() {
			path = path + "/accounts"
		})

		It("should return a POST request", func() {
			type payload struct {
				Data *account.Account `json:"data"`
			}

			attrs := account.NewAttributes("EUR", "FR")
			acc := account.New(orgID).WithID(accountID).WithAttributes(attrs)

			req, err := cl.NewRequest(ctx, http.MethodPost, path, nil, payload{Data: acc})
			Expect(err).To(BeNil())

			By("setting the correct request URL", func() {
				Expect(req.URL.String()).To(Equal("https://" + host + path))
			})

			By("setting the correct HTTP method", func() {
				Expect(req.Method).To(Equal(http.MethodPost))
			})

			By("including a request body", func() {
				Expect(req.Body).To(Not(BeNil()))

				var p payload
				err := json.NewDecoder(req.Body).Decode(&p)
				Expect(err).To(BeNil())
				Expect(p.Data).To(BeEquivalentTo(acc))
			})

			By("setting the default headers", func() {
				Expect(req.Header.Get("Host")).To(Equal(host))
				Expect(req.Header.Get("Accept")).To(Equal("application/vnd.api+json"))
				Expect(req.Header.Get("Date")).To(Not(BeEmpty()))
				Expect(req.Header.Get("User-Agent")).To(Not(BeEmpty()))
			})

			By("setting the content type header", func() {
				Expect(req.Header.Get("Content-Type")).To(Equal("application/vnd.api+json"))
			})
		})
	})

	Describe("Creating a DELETE request", func() {
		BeforeEach(func() {
			path = path + "/accounts/" + accountID
		})

		It("should return a DELETE request", func() {
			params := map[string]string{"version": "1"}

			req, err := cl.NewRequest(ctx, http.MethodDelete, path, params, nil)
			Expect(err).To(BeNil())

			By("setting the correct request URL", func() {
				Expect(req.URL.String()).To(Equal("https://" + host + path + "?version=1"))
			})

			By("setting the correct HTTP method", func() {
				Expect(req.Method).To(Equal(http.MethodDelete))
			})

			By("not including a request body", func() {
				Expect(req.Body).To(BeNil())
			})

			By("setting the default headers", func() {
				Expect(req.Header.Get("Host")).To(Equal(host))
				Expect(req.Header.Get("Accept")).To(Equal("application/vnd.api+json"))
				Expect(req.Header.Get("Date")).To(Not(BeEmpty()))
				Expect(req.Header.Get("User-Agent")).To(Not(BeEmpty()))
			})
		})
	})

	Describe("Decoding errors", func() {
		BeforeEach(func() {
			path = path + "/accounts/" + accountID
		})

		Context("bad request error", func() {
			BeforeEach(func() {
				fakeHTTPClient.DoReturns(&http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       io.NopCloser(bytes.NewBuffer([]byte(fixtures.BadRequestError))),
					Request: &http.Request{
						Method: http.MethodGet,
						URL:    &url.URL{},
					},
				}, nil)
			})

			It("should decode the HTTP response", func() {
				var b []byte
				resp, err := cl.Get(ctx, path, nil, &b)
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())

				assertErrorResponse(err, http.StatusBadRequest)
			})
		})

		Context("conflict error", func() {
			BeforeEach(func() {
				fakeHTTPClient.DoReturns(&http.Response{
					StatusCode: http.StatusConflict,
					Body:       io.NopCloser(bytes.NewBuffer([]byte(fixtures.ConflictError))),
					Request: &http.Request{
						Method: http.MethodPost,
						URL:    &url.URL{},
					},
				}, nil)
			})

			It("should decode the HTTP response", func() {
				params := map[string]string{"version": "1"}
				resp, err := cl.Delete(ctx, path, params)
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())

				assertErrorResponse(err, http.StatusConflict)
			})
		})

		Context("forbidden error", func() {
			BeforeEach(func() {
				fakeHTTPClient.DoReturns(&http.Response{
					StatusCode: http.StatusForbidden,
					Body:       io.NopCloser(bytes.NewBuffer([]byte(fixtures.ForbiddenError))),
					Request: &http.Request{
						Method: http.MethodGet,
						URL:    &url.URL{},
					},
				}, nil)
			})

			It("should decode the HTTP response", func() {
				var b []byte
				resp, err := cl.Get(ctx, path, nil, &b)
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())

				assertErrorResponse(err, http.StatusForbidden)
			})
		})

		Context("other errors", func() {
			BeforeEach(func() {
				fakeHTTPClient.DoReturns(&http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       http.NoBody,
					Request: &http.Request{
						Method: http.MethodGet,
						URL:    &url.URL{},
					},
				}, nil)
			})

			It("should decode the HTTP response", func() {
				var b []byte
				resp, err := cl.Get(ctx, path, nil, &b)
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())

				assertErrorResponse(err, http.StatusInternalServerError)
			})
		})
	})

	Describe("Executing GET requests", func() {
		var respBody string

		BeforeEach(func() {
			respBody = fixtures.AccountsResponseAllFields(orgID, accountID, "FR", "EUR")
		})

		Context("the request is successful", func() {
			BeforeEach(func() {
				fakeHTTPClient.DoReturns(&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBuffer([]byte(respBody))),
				}, nil)
			})

			It("should return the HTTP success response", func() {
				var r response
				resp, err := cl.Get(ctx, path, nil, &r)
				Expect(err).To(BeNil())

				Expect(resp.StatusCode).To(Equal(http.StatusOK))
				expectAccountsResponse(r, orgID, accountID)
			})
		})

		Context("the request fails", func() {
			BeforeEach(func() {
				fakeHTTPClient.DoReturns(nil, fmt.Errorf("request error"))
			})

			It("should return the error", func() {
				var r response
				resp, err := cl.Get(ctx, path, nil, &r)
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Executing POST requests", func() {
		var (
			respBody string

			attrs *account.Attributes
			acc   *account.Account
		)

		BeforeEach(func() {
			path = path + "/accounts/" + accountID

			attrs = account.NewAttributes("EUR", "FR")
			acc = account.New(orgID).WithID(accountID).WithAttributes(attrs)

			respBody = fixtures.AccountsResponseAllFields(orgID, accountID, "FR", "EUR")
		})

		Context("the request is successful", func() {
			BeforeEach(func() {
				fakeHTTPClient.DoReturns(&http.Response{
					StatusCode: http.StatusCreated,
					Body:       io.NopCloser(bytes.NewBuffer([]byte(respBody))),
				}, nil)
			})

			It("should return the HTTP success response", func() {
				var r response
				resp, err := cl.Post(ctx, path, acc, &r)
				Expect(err).To(BeNil())
				Expect(resp).To(Not(BeNil()))

				Expect(resp.StatusCode).To(Equal(http.StatusCreated))
				expectAccountsResponse(r, orgID, accountID)
			})
		})

		Context("the request fails", func() {
			BeforeEach(func() {
				fakeHTTPClient.DoReturns(nil, fmt.Errorf("request error"))
			})

			It("should return the error", func() {
				var body []byte
				resp, err := cl.Post(ctx, path, acc, &body)
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Executing DELETE requests", func() {
		BeforeEach(func() {
			path = path + "/accounts/" + accountID
		})

		Context("the request is successful", func() {
			BeforeEach(func() {
				fakeHTTPClient.DoReturns(&http.Response{
					StatusCode: http.StatusNoContent,
					Body:       http.NoBody,
				}, nil)
			})

			It("should return the HTTP success response", func() {
				params := map[string]string{"version": "1"}
				resp, err := cl.Delete(ctx, path, params)
				Expect(err).To(BeNil())
				Expect(resp).To(Not(BeNil()))

				Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
			})
		})

		Context("the request fails", func() {
			BeforeEach(func() {
				fakeHTTPClient.DoReturns(nil, fmt.Errorf("request error"))
			})

			It("should return the error", func() {
				resp, err := cl.Delete(ctx, path, nil)
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())
			})
		})
	})
})

type response struct {
	Data  *account.Account `json:"data"`
	Links *account.Links   `json:"links"`
}

func expectAccountsResponse(r response, orgID string, accountID string) {
	Expect(r.Data).To(Not(BeNil()))
	Expect(r.Data.ID).To(Equal(accountID))
	Expect(r.Data.OrganisationID).To(Equal(orgID))

	Expect(r.Links).To(Not(BeNil()))
}

type errResponse interface {
	HTTPResponse() *http.Response
	Error() string
}

func assertErrorResponse(err error, statusCode int) {
	var e errResponse
	Expect(errors.As(err, &e)).To(BeTrue())
	Expect(e.HTTPResponse().StatusCode).To(Equal(statusCode))
}
