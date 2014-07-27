package heiankyoview_test

import (
	. "github.com/gyuho/goraph/viz/heiankyoview"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinarySearch", func() {
	It("Test", func() {
		arr := []int{1, 3, 6, 10}
		Expect(BinarySearch(arr, 1)).To(Equal(0))
		Expect(BinarySearch(arr, 2)).To(Equal(-1))
	})
})

var _ = Describe("Coordinate", func() {
	var co *Coordinate
	BeforeEach(func() {
		co = NewCoordinate(0, 3)
		co.Insert(2, 10)
		co.Insert(0, -5)
	})

	It("returns state", func() {
		Expect(co.Size()).To(Equal(4))
		Expect(co.MinLine()).To(Equal(-5))
		Expect(co.MaxLine()).To(Equal(10))
		Expect(co.Width()).To(Equal(15))
		Expect(co.IndexOf(3)).To(Equal(2))
	})

	It("calcs left intersection", func() {
		var cases = []struct {
			index int
			line  int
			from  int
			to    int
		}{
			{2, 1, 1, 1},
			{2, 0, 1, 1},
			{2, -3, 0, 1},
			{2, -5, 0, 1},
			{2, -7, -1, 1},
		}
		for _, c := range cases {
			_f, _t := co.GetLeftIntersection(c.index, c.line)
			Expect(_f).To(Equal(c.from))
			Expect(_t).To(Equal(c.to))
		}
	})

	It("calcs right intersection", func() {
		var cases = []struct {
			index int
			line  int
			from  int
			to    int
		}{
			{1, 1, 1, 1},
			{1, 3, 1, 1},
			{1, 5, 0, 1},
			{1, 10, 1, 2},
			{1, 11, 1, 3},
		}
		for _, c := range cases {
			_f, _t := co.GetRightIntersection(c.index, c.line)
			Expect(_f).To(Equal(c.from))
			Expect(_t).To(Equal(c.to))
		}

	})
})
