package integration_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var baseURL string

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = BeforeSuite(func() {
	baseURL = os.Getenv("ACCOUNTS_API_BASE_URL")
	if baseURL == "" {
		Skip("ACCOUNTS_API_BASE_URL env var is not set")
		return
	}
})
