package heiankyoview

import "sort"

type Val int   // Coordinate Value
type Index int // Index

type Rectangle struct {
	x Val
	y Val
	w Val
	h Val
}

func mkR(w, h Val) *Rectangle {
	return &Rectangle{w: w, h: h}
}

type Packing struct {
}

func (p *Packing) add(r *Rectangle) {
}

func BinarySearch(xs []Val, n Val) Index {
	i := sort.Search(len(xs), func(i int) bool { return xs[i] >= n })
	if i < len(xs) && xs[i] == n {
		return Index(i)
	} else {
		return Index(-1)
	}
}

type Coordinate struct {
	xs []Val
}

func NewCoordinate(minLine, maxLine Val) *Coordinate {
	c := Coordinate{[]Val{minLine, maxLine}}
	return &c
}

func (c *Coordinate) get(n Index) Val {
	return c.xs[n]
}

func (c *Coordinate) has(line Val) bool {
	return false // TODO
}

// Insertion is costful but there's no way but this
// It only happens per adding a rectangle thus not frequent
func (c *Coordinate) insert(at Index, line Val) {
	before := c.xs[:at]
	before = append(before, line)
	after := c.xs[at:]
	c.xs = append(before, after...) // I don't know how efficient this is
}

func (c *Coordinate) IndexOf(line Val) Index {
	return BinarySearch(c.xs, line)
}

func (c *Coordinate) minLine() Val {
	return c.xs[0]
}

func (c *Coordinate) maxLine() Val {
	return c.xs[len(c.xs)-1]
}

func (c *Coordinate) width() Val {
	return c.maxLine() - c.minLine()
}

func (c *Coordinate) getRightIntersection(i Index, line Val) (Index, Index) {
	return 0, 0 // TODO
}

func (c *Coordinate) getLeftIntersection(i Index, line Val) (Index, Index) {
	return 0, 0 // TODO
}
