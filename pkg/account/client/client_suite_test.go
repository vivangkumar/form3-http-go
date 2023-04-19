package client_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAccountClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Accounts Client Suite")
}
