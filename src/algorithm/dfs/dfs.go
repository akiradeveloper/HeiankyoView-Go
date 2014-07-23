// Package dfs implements Dapth First Search algorithm (or DFS).
package dfs

/*
Pseudo Code CLRS page 604

DFS(G)
	for each vertex u ∈ G.V
		u.color = WHITE
		u.π = NIL
	// this is already done
	// when instantiating the graph
	// and instead of InVertices
	// we can just create another slice
	// inside Graph (Prev)
	// in order not to modify the original graph

	time = 0

	for each vertex u ∈ G.V
		if u.color == WHITE
			DFS-Visit(G, u)

---

DFS-Visit(G, u)
	time = time + 1
	u.d = time
	u.color = GRAY

	for each v ∈ G.Adj[u]
		if v.color == WHITE
			v.π = u
			DFS-Visit(G, v)

	u.color = BLACK
	time = time + 1
	u.f = time
*/

import (
	"fmt"

	"github.com/gyuho/goraph/graph/gsd"
	slice "github.com/gyuho/gosequence"
)

// Breadth First Search: Queue
// Depth First Search: Stack, Recursion

// DFS does Depth First Search and return the result in visited order.
func DFS(g *gsd.Graph) string {
	// time = 0
	var stamp int64 = 0

	// for each vertex u ∈ g.V
	result := slice.NewSequence()
	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		if vtx.(*gsd.Vertex).Color == "white" {
			DFSVisit(g, vtx.(*gsd.Vertex), stamp, result)
		}
	}

	s := ""
	for _, v := range *result {
		s += fmt.Sprintf("%v → ", v)
	}
	return s[:len(s)-5]
}

// DFSVisit recursively visits the vertices in the graph.
func DFSVisit(g *gsd.Graph, src *gsd.Vertex, stamp int64, result *slice.Sequence) {
	stamp += 1
	src.StampD = stamp
	src.Color = "gray"

	ovs := src.GetOutVertices()
	for _, vtx := range *ovs {
		if vtx == nil {
			continue
		}
		if vtx.(*gsd.Vertex).Color == "white" {
			vtx.(*gsd.Vertex).Prev.PushBack(src)
			DFSVisit(g, vtx.(*gsd.Vertex), stamp, result)
		}
	}

	src.Color = "black"
	result.PushFront(src.ID)

	stamp += 1
	src.StampF = stamp
}

// Path returns true if there is a path between two Vertices.
func Path(g *gsd.Graph, src, end *gsd.Vertex) bool {
	if src == nil || end == nil {
		panic("Wrong Vertex Passed!")
	}
	rb := false
	if src.Color == "white" {
		path(g, src, end, &rb)
	}
	return rb
}

// path recursively visits the vertices in the graph.
func path(g *gsd.Graph, src, end *gsd.Vertex, rb *bool) {
	if src == end {
		*rb = true
		return
	}
	src.Color = "gray"
	ovs := src.GetOutVertices()
	for _, vtx := range *ovs {
		if vtx == nil {
			continue
		}
		if vtx.(*gsd.Vertex).Color == "white" {
			path(g, vtx.(*gsd.Vertex), end, rb)
		}
	}
	src.Color = "black"
}
