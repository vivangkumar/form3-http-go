package integration_test

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vivangkumar/form3-http-go/pkg/account"
	"github.com/vivangkumar/form3-http-go/pkg/client"
	"github.com/vivangkumar/form3-http-go/pkg/form3"
)

var _ = Describe("Accounts API", func() {
	var (
		ctx context.Context

		orgID     string
		accountID string

		acc   *account.Account
		attrs *account.Attributes

		cl *form3.Client
	)

	BeforeEach(func() {
		ctx = context.Background()

		orgID = uuid.NewString()
		accountID = uuid.NewString()

		var err error
		cl, err = form3.New(client.WithBaseURL(baseURL))
		Expect(err).To(BeNil())

		attrs = account.NewAttributes("EUR", "FR").
			WithBankID("20041").
			WithBankIDCode("FR").
			WithAccountNumber("0500013M026").
			WithIban("FR1420041010050500013M02606").
			WithBic("NWBKFR42").
			WithName("eur-fr-bank-acc")

		acc = account.
			New(orgID).
			WithID(accountID).
			WithAttributes(attrs)
	})

	Describe("Creating a new account", func() {
		When("a valid request is made", func() {
			It("should return the created account", func() {
				resp, err := cl.Accounts.Create(ctx, acc)
				Expect(err).To(BeNil())

				created := resp.Data
				Expect(created.ID).To(Equal(accountID))
				Expect(created.OrganisationID).To(Equal(orgID))
				Expect(created.Type).To(Equal("accounts"))
				Expect(created.Version).To(Not(BeNil()))

				Expect(resp.Links).To(Not(BeNil()))
				Expect(resp.Links.Self).To(Not(BeNil()))
			})
		})

		When("an invalid request is made", func() {
			BeforeEach(func() {
				acc.OrganisationID = "invalid-org-id"
			})

			It("should return a bad request error", func() {
				resp, err := cl.Accounts.Create(ctx, acc)
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())

				var e errResponse
				Expect(errors.As(err, &e)).To(BeTrue())
				Expect(e.HTTPResponse().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(e.Error()).To(Not(BeEmpty()))
			})
		})
	})

	Context("An account has been created", func() {
		var created *account.Account

		BeforeEach(func() {
			resp, err := cl.Accounts.Create(ctx, acc)
			Expect(err).To(BeNil())

			created = resp.Data
		})

		When("fetching the account", func() {
			It("should return the created account", func() {
				resp, err := cl.Accounts.Fetch(ctx, account.FetchAccountParams{
					ID: created.ID,
				})
				Expect(err).To(BeNil())
				Expect(resp.Data).To(BeEquivalentTo(created))
				Expect(resp.Links).To(Not(BeNil()))
				Expect(resp.Links.Self).To(Not(BeEmpty()))
			})
		})

		When("the account has been deleted", func() {
			BeforeEach(func() {
				resp, err := cl.Accounts.Delete(ctx, account.DeleteAccountParams{
					ID:      created.ID,
					Version: *created.Version,
				})
				Expect(err).To(BeNil())
				Expect(resp).To(Not(BeNil()))
			})

			Context("and the account is fetched", func() {
				It("should return a not found error", func() {
					resp, err := cl.Accounts.Fetch(ctx, account.FetchAccountParams{
						ID: created.ID,
					})
					Expect(err).To(Not(BeNil()))
					Expect(resp).To(BeNil())

					var e errResponse
					Expect(errors.As(err, &e)).To(BeTrue())
					Expect(e.HTTPResponse().StatusCode).To(Equal(http.StatusNotFound))
					Expect(e.Error()).To(Not(BeEmpty()))
				})
			})
		})
	})

	Describe("Deleting an account", func() {
		var created *account.Account

		BeforeEach(func() {
			resp, err := cl.Accounts.Create(ctx, acc)
			Expect(err).To(BeNil())

			created = resp.Data
			Expect(created.Version).To(Not(BeNil()))
		})

		When("a valid request is made", func() {
			It("should delete the account", func() {
				resp, err := cl.Accounts.Delete(ctx, account.DeleteAccountParams{
					ID:      created.ID,
					Version: *created.Version,
				})
				Expect(err).To(BeNil())
				Expect(resp).To(Not(BeNil()))
			})
		})

		When("an invalid account ID is passed", func() {
			It("should return a bad request error", func() {
				resp, err := cl.Accounts.Delete(ctx, account.DeleteAccountParams{
					ID:      "some-wrong-id",
					Version: *created.Version,
				})
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())

				var e errResponse
				Expect(errors.As(err, &e)).To(BeTrue())
				Expect(e.HTTPResponse().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(e.Error()).To(Not(BeEmpty()))
			})
		})

		When("a wrong version is passed", func() {
			It("should return a conflict error", func() {
				resp, err := cl.Accounts.Delete(ctx, account.DeleteAccountParams{
					ID:      created.ID,
					Version: 1,
				})
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())

				var e errResponse
				Expect(errors.As(err, &e)).To(BeTrue())
				Expect(e.HTTPResponse().StatusCode).To(Equal(http.StatusConflict))
				Expect(e.Error()).To(Not(BeEmpty()))
			})
		})

		When("an unknown account ID is passed", func() {
			It("should return a not found error", func() {
				resp, err := cl.Accounts.Delete(ctx, account.DeleteAccountParams{
					ID:      uuid.NewString(),
					Version: *created.Version,
				})
				Expect(err).To(Not(BeNil()))
				Expect(resp).To(BeNil())

				var e errResponse
				Expect(errors.As(err, &e)).To(BeTrue())
				Expect(e.HTTPResponse().StatusCode).To(Equal(http.StatusNotFound))
				Expect(e.Error()).To(Not(BeEmpty()))
			})
		})
	})
})

type errResponse interface {
	HTTPResponse() *http.Response
	Error() string
}
