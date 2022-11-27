package drmaa2interface_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDrmaa2interface(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Drmaa2interface Suite")
}
