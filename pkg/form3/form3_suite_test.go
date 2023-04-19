package form3_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestForm3Client(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Form3 Client Suite")
}
