package heiankyoview

import "sort"

type val int   // Coordinate value
type index int // Index

type Rectangle struct {
	x val
	y val
	w val
	h val
}

type Packing struct {
}

func (p *Packing) add(r *Rectangle) {
}

func binarySearch(xs []val, n val) index {
	i := sort.Search(len(xs), func(i int) bool { return xs[i] >= n })
	if i < len(xs) && xs[i] == n {
		return index(i)
	} else {
		return index(-1)
	}
}

type Coordinate struct {
	xs []val
}

func NewCoordinate(minLine, maxLine val) *Coordinate {
	c := Coordinate{[]val{minLine, maxLine}}
	return &c
}

func (c *Coordinate) get(n index) val {
	return c.xs[n]
}

func (c *Coordinate) has(line val) bool {
	return false // TODO
}

// Insertion is costful but there's no way but this
// It only happens per adding a rectangle thus not frequent
func (c *Coordinate) insert(at index, line val) {
	before := c.xs[:at]
	before = append(before, line)
	after := c.xs[at:]
	c.xs = append(before, after...) // I don't know how efficient this is
}

func (c *Coordinate) indexOf(line val) index {
	return binarySearch(c.xs, line)
}

func (c *Coordinate) minLine() val {
	return c.xs[0]
}

func (c *Coordinate) maxLine() val {
	return c.xs[len(c.xs)-1]
}

func (c *Coordinate) width() val {
	return c.maxLine() - c.minLine()
}

func (c *Coordinate) getRightIntersection(i index, line val) (index, index) {
	return 0, 0 // TODO
}

func (c *Coordinate) getLeftIntersection(i index, line val) (index, index) {
	return 0, 0 // TODO
}
