package form3_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vivangkumar/form3-http-go/pkg/client"
	"github.com/vivangkumar/form3-http-go/pkg/form3"
)

var _ = Describe("Form3 client", func() {
	Describe("Client creation", func() {
		It("should create a client without opts", func() {
			client, err := form3.New()
			Expect(err).To(BeNil())
			Expect(client).To(Not(BeNil()))
		})

		It("should create a client with opts", func() {
			client, err := form3.New(
				client.WithBaseURL("http://localhost:8080"),
			)

			Expect(err).To(BeNil())
			Expect(client).To(Not(BeNil()))
		})
	})
})
