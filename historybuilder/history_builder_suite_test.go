package historybuilder_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHistoryBuilder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HistoryBuilder Suite")
}
