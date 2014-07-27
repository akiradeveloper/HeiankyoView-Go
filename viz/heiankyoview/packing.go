package heiankyoview

import "sort"

type Rectangle struct {
	x int
	y int
	w int
	h int
}

func mkR(w, h int) *Rectangle {
	return &Rectangle{w: w, h: h}
}

type Packing struct {
}

func (p *Packing) add(r *Rectangle) {
}

func BinarySearch(xs []int, n int) int {
	i := sort.Search(len(xs), func(i int) bool { return xs[i] >= n })
	if i < len(xs) && xs[i] == n {
		return i
	} else {
		return -1
	}
}

type Coordinate struct {
	xs []int
}

func NewCoordinate(minLine, maxLine int) *Coordinate {
	c := Coordinate{[]int{minLine, maxLine}}
	return &c
}

func (c *Coordinate) Size() int {
	return len(c.xs)
}

func (c *Coordinate) Get(n int) int {
	return c.xs[n]
}

func (c *Coordinate) Has(line int) bool {
	return false
}

// Insertion is costful but there's no way but this
// It only happens per adding a rectangle thus not frequent
func (c *Coordinate) Insert(at int, line int) {
	before := c.xs[:at]
	before = append(before, line)
	after := c.xs[at:]
	c.xs = append(before, after...) // I don't know how efficient this is
}

func (c *Coordinate) IndexOf(line int) int {
	return BinarySearch(c.xs, line)
}

func (c *Coordinate) MinLine() int {
	return c.xs[0]
}

func (c *Coordinate) MaxLine() int {
	return c.xs[len(c.xs)-1]
}

func (c *Coordinate) Width() int {
	return c.MaxLine() - c.MinLine()
}

func (c *Coordinate) GetRightIntersection(i int, line int) (int, int) {
	return 0, 0 // TODO
}

func (c *Coordinate) GetLeftIntersection(i int, line int) (int, int) {
	return 0, 0 // TODO
}
