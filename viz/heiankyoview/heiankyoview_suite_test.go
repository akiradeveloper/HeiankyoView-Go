package heiankyoview_test

import (
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var _ = Describe("Testing Sort library", func() {
	It("Search", func() {
		arr := []int{1, 3, 6, 10}
		// TODO table-driven
		Expect(sort.Search(len(arr), func(i int) bool { return arr[i] >= 1 })).To(Equal(0))
		Expect(sort.Search(len(arr), func(i int) bool { return arr[i] >= 2 })).To(Equal(1))
		Expect(sort.Search(len(arr), func(i int) bool { return arr[i] >= 3 })).To(Equal(1))
	})

	// SearchInts - Find the index to best insert the element
	It("SearchInts", func() {
		arr := []int{1, 3}
		// TODO table-driven
		Expect(sort.SearchInts(arr, 0)).To(Equal(0))
		Expect(sort.SearchInts(arr, 1)).To(Equal(0))
		Expect(sort.SearchInts(arr, 2)).To(Equal(1))
		Expect(sort.SearchInts(arr, 3)).To(Equal(1))
	})
})

func TestHeiankyoview(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Heiankyoview Suite")
}
