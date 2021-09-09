package verus_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVerus(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Verus Suite")
}
