package integration_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
	"testing"
)

var deleter string

func TestIntegrationTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IntegrationTests Suite")
}

var _ = BeforeSuite(func() {
	var err error
	deleter, err = gexec.Build("github.com/aclevername/concourse-flake-detector/cmd/flake-detector/")
	Expect(err).NotTo(HaveOccurred())
})
