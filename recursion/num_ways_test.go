package recursion

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCurrency(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Currency Suite")
}

var _ = Describe("Currency", func() {
	It("should handle American coins", func() {
		// denominations are in cents and they should be sorted in descending order
		curr := NewCurrency([]int{100, 50, 25, 10, 5, 1})
		Expect(curr.NumWays(100)).To(Equal(293))
		Expect(curr.MinChange(7)).To(Equal(3))
		Expect(curr.MinChange(194)).To(Equal(9))
	})
})
