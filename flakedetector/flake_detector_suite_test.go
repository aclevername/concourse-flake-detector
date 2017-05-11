package flakedetector_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestComparator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Comparator Suite")
}
