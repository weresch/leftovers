package groupingobjects_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGroupingObjects(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "nsxt/groupingobjects")
}
