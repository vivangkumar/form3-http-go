package client_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vivangkumar/form3-http-go/pkg/account"
	"github.com/vivangkumar/form3-http-go/pkg/account/client"
	"github.com/vivangkumar/form3-http-go/pkg/account/internal/fakes"
	"github.com/vivangkumar/form3-http-go/pkg/internal/fixtures"
)

var _ = Describe("Account client", func() {
	var (
		ctx context.Context

		orgID     string
		accountID string

		fakeBaseClient *fakes.FakeBaseClient
		cl             *client.Client

		respBody string
		acc      *account.Account
		attrs    *account.Attributes
	)

	BeforeEach(func() {
		ctx = context.Background()

		orgID = uuid.NewString()
		accountID = uuid.NewString()

		fakeBaseClient = new(fakes.FakeBaseClient)
		cl = client.New(fakeBaseClient)

		attrs = account.NewAttributes("EUR", "FR")
		acc = account.New(orgID).WithID(accountID).WithAttributes(attrs)
	})

	Describe("Create account", func() {
		Context("with success response", func() {
			BeforeEach(func() {
				fakeBaseClient.PostStub = func(
					ctx context.Context,
					path string,
					body any,
					target any,
				) (*http.Response, error) {
					err := json.Unmarshal([]byte(respBody), &target)
					if err != nil {
						return nil, err
					}

					return &http.Response{
						StatusCode: http.StatusCreated,
						Body:       io.NopCloser(bytes.NewBuffer([]byte(respBody))),
					}, nil
				}
			})

			When("all fields are returned", func() {
				BeforeEach(func() {
					respBody = fixtures.AccountsResponseAllFields(orgID, accountID, "FR", "EUR")
				})

				It("should return the serialised account data and links", func() {
					resp, err := cl.Create(ctx, acc)
					Expect(err).To(BeNil())
					Expect(resp).To(Not(BeNil()))

					assertAllAccountFields(resp.Data, orgID, accountID)
					assertAllLinkFields(resp.Links)
				})
			})

			When("only the minimum fields are returned", func() {
				BeforeEach(func() {
					respBody = fixtures.AccountsResponseMinFields(orgID, accountID, "FR", "EUR")
				})

				It("should return the account data and links", func() {
					resp, err := cl.Create(ctx, acc)
					Expect(err).To(BeNil())
					Expect(resp).To(Not(BeNil()))

					assertMinAccountFields(resp.Data, orgID, accountID)
					assertMinLinkFields(resp.Links)
				})
			})
		})

		Context("with empty account entity", func() {
			It("should return an error", func() {
				resp, err := cl.Create(ctx, nil)
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())
			})
		})

		Context("with request error", func() {
			BeforeEach(func() {
				fakeBaseClient.PostReturns(nil, fmt.Errorf("request error"))
			})

			It("should return an error", func() {
				resp, err := cl.Create(ctx, acc)
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())
			})
		})

		Context("with HTTP error response", func() {
			errMsg := "BAD_REQUEST"
			BeforeEach(func() {
				fakeBaseClient.PostReturns(nil, apiError{
					httpResponse: &http.Response{
						StatusCode: http.StatusBadRequest,
					},
					underlying: fmt.Errorf(errMsg),
				})
			})

			It("should return an http response enriched error", func() {
				resp, err := cl.Create(ctx, acc)
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())

				assertErrorResponse(err, http.StatusBadRequest, errMsg)
			})
		})

	})

	Describe("Fetch account", func() {
		Context("with success response", func() {
			BeforeEach(func() {
				fakeBaseClient.GetStub = func(
					ctx context.Context,
					path string,
					query map[string]string,
					target any,
				) (*http.Response, error) {
					err := json.Unmarshal([]byte(respBody), &target)
					if err != nil {
						return nil, err
					}

					return &http.Response{
						StatusCode: http.StatusCreated,
						Body:       io.NopCloser(bytes.NewBuffer([]byte(respBody))),
					}, nil
				}
			})

			When("all fields are returned", func() {
				BeforeEach(func() {
					respBody = fixtures.AccountsResponseAllFields(orgID, accountID, "FR", "EUR")
				})

				It("should return the serialised account data and links", func() {
					resp, err := cl.Fetch(ctx, account.FetchAccountParams{
						ID: accountID,
					})
					Expect(err).To(BeNil())
					Expect(resp).To(Not(BeNil()))

					assertAllAccountFields(resp.Data, orgID, accountID)
					assertAllLinkFields(resp.Links)
				})
			})

			When("only the minimum fields are returned", func() {
				BeforeEach(func() {
					respBody = fixtures.AccountsResponseMinFields(orgID, accountID, "FR", "EUR")
				})

				It("should return the account data and links", func() {
					resp, err := cl.Fetch(ctx, account.FetchAccountParams{
						ID: accountID,
					})
					Expect(err).To(BeNil())
					Expect(resp).To(Not(BeNil()))

					assertMinAccountFields(resp.Data, orgID, accountID)
					assertMinLinkFields(resp.Links)
				})
			})
		})

		Context("with request error", func() {
			BeforeEach(func() {
				fakeBaseClient.GetReturns(nil, fmt.Errorf("request error"))
			})

			It("should return an error", func() {
				resp, err := cl.Fetch(ctx, account.FetchAccountParams{
					ID: accountID,
				})
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())
			})
		})

		Context("with HTTP error response", func() {
			errMsg := "NOT_FOUND"
			BeforeEach(func() {
				fakeBaseClient.GetReturns(nil, apiError{
					httpResponse: &http.Response{
						StatusCode: http.StatusNotFound,
					},
					underlying: fmt.Errorf(errMsg),
				})
			})

			It("should return an http response enriched error", func() {
				resp, err := cl.Fetch(ctx, account.FetchAccountParams{
					ID: accountID,
				})
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())

				assertErrorResponse(err, http.StatusNotFound, errMsg)
			})
		})
	})

	Describe("Delete account", func() {
		Context("with success response", func() {
			BeforeEach(func() {
				fakeBaseClient.DeleteReturns(&http.Response{
					StatusCode: http.StatusNoContent,
				}, nil)
			})

			It("should return a response", func() {
				resp, err := cl.Delete(ctx, account.DeleteAccountParams{
					ID:      accountID,
					Version: 1,
				})
				Expect(err).To(BeNil())
				Expect(resp).To(Not(BeNil()))
			})
		})

		Context("with request error", func() {
			BeforeEach(func() {
				fakeBaseClient.DeleteReturns(nil, fmt.Errorf("request error"))
			})

			It("should return an error", func() {
				resp, err := cl.Delete(ctx, account.DeleteAccountParams{
					ID:      accountID,
					Version: 1,
				})
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())
			})
		})

		Context("with HTTP error response", func() {
			errMsg := "CONFLICT"
			BeforeEach(func() {
				fakeBaseClient.DeleteReturns(nil, apiError{
					httpResponse: &http.Response{
						StatusCode: http.StatusConflict,
					},
					underlying: fmt.Errorf(errMsg),
				})
			})

			It("should return an http response enriched error", func() {
				resp, err := cl.Delete(ctx, account.DeleteAccountParams{
					ID:      accountID,
					Version: 1,
				})
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())

				assertErrorResponse(err, http.StatusConflict, errMsg)
			})
		})
	})
})

type apiError struct {
	httpResponse *http.Response
	underlying   error
}

func (a apiError) HTTPResponse() *http.Response {
	return a.httpResponse
}

func (a apiError) Error() string {
	return fmt.Sprintf(a.underlying.Error())
}

type errResponse interface {
	HTTPResponse() *http.Response
	Error() string
}

func assertErrorResponse(err error, statusCode int, msg string) {
	var e errResponse
	Expect(errors.As(err, &e)).To(BeTrue())
	Expect(e.HTTPResponse().StatusCode).To(Equal(statusCode))
	Expect(e.Error()).To(Equal(msg))
}

func assertAllAccountFields(acc *account.Account, orgID string, accountID string) {
	assertCommonAccountFields(acc, orgID, accountID)

	Expect(acc.Attributes.BankID).To(Not(BeEmpty()))
	Expect(acc.Attributes.BankIDCode).To(Not(BeEmpty()))
	Expect(acc.Attributes.Bic).To(Not(BeEmpty()))
	Expect(acc.Attributes.CustomerID).To(Not(BeNil()))
}

func assertCommonAccountFields(acc *account.Account, orgID string, accountID string) {
	Expect(acc.ID).To(Equal(accountID))
	Expect(acc.OrganisationID).To(Equal(orgID))
	Expect(acc.Attributes.AccountClassification).To(Not(BeNil()))
	Expect(acc.Attributes.AccountMatchingOptOut).To(Not(BeNil()))
	Expect(acc.Attributes.AccountNumber).To(Not(BeEmpty()))
	Expect(acc.Attributes.BaseCurrency).To(Not(BeEmpty()))
	Expect(acc.Attributes.Iban).To(Not(BeEmpty()))
	Expect(acc.Attributes.Country).To(Not(BeEmpty()))
	Expect(acc.Attributes.JointAccount).To(Not(BeNil()))
	Expect(acc.Attributes.Status).To(Not(BeNil()))
	Expect(acc.Attributes.Switched).To(Not(BeNil()))
}

func assertMinAccountFields(acc *account.Account, orgID string, accountID string) {
	assertCommonAccountFields(acc, orgID, accountID)

	Expect(acc.Attributes.BankID).To(BeEmpty())
	Expect(acc.Attributes.BankIDCode).To(BeEmpty())
	Expect(acc.Attributes.Bic).To(BeEmpty())
	Expect(acc.Attributes.CustomerID).To(BeNil())
}

func assertAllLinkFields(links *account.Links) {
	Expect(links.Self).To(Not(BeNil()))
	Expect(links.First).To(Not(BeNil()))
	Expect(links.Last).To(Not(BeNil()))
	Expect(links.Next).To(Not(BeNil()))
	Expect(links.Prev).To(Not(BeNil()))
}

func assertMinLinkFields(links *account.Links) {
	Expect(links.Self).To(Not(BeNil()))
	Expect(links.First).To(BeNil())
	Expect(links.Last).To(BeNil())
	Expect(links.Next).To(BeNil())
	Expect(links.Prev).To(BeNil())
}
