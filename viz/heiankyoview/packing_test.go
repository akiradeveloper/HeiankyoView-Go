package heiankyoview

import (
	"sort"
	"testing"
)

func mkR(w, h val) *Rectangle {
	return &Rectangle{w: w, h: h}
}

func assertR(t *testing.T, x, y val, r *Rectangle) {
	if r.x != x || r.y != y {
		t.Errorf("expected x(%d) y(%d) but actual x(%d) y(%d)", x, y, r.x, r.y)
	}
}

func TestTest(t *testing.T) {
	// Test assertR
	assertR(t, 0, 0, mkR(2, 2))
}

func TestSearchInts(t *testing.T) {
	arr := []int{1, 3}

	// SearchInts - Find the index to best insert the element

	if sort.SearchInts(arr, 0) != 0 {
		t.Fail()
	}
	if sort.SearchInts(arr, 1) != 0 {
		t.Fail()
	}
	if sort.SearchInts(arr, 2) != 1 {
		t.Fail()
	}
	if sort.SearchInts(arr, 3) != 1 {
		t.Fail()
	}
}

func TestSearch(t *testing.T) {
	arr := []int{1, 3, 6, 10}

	// Search - Find the index where the given predicate first turns true
	// using binary search algorithm.

	if sort.Search(len(arr), func(i int) bool { return arr[i] >= 1 }) != 0 {
		t.Fail()
	}

	if sort.Search(len(arr), func(i int) bool { return arr[i] >= 3 }) != 1 {
		t.Fail()
	}

	if sort.Search(len(arr), func(i int) bool { return arr[i] >= 2 }) != 1 {
		t.Fail()
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []val{1, 3, 6, 10}
	if binarySearch(arr, 1) != 0 {
		t.Fail()
	}
	if binarySearch(arr, 2) != -1 {
		t.Fail()
	}
}
