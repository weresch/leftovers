package logicalrouting

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLogicalRouting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "nsxt/logicalrouting")
}
