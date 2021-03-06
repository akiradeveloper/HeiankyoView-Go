package example_with_testing

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/algorithm/spbf"
	"github.com/gyuho/goraph/graph/gsd"
	// go test -v github.com/gyuho/goraph/example_with_testing
	// go test -v /Users/gyuho/go/src/github.com/gyuho/goraph/example_with_testing/spbf_json_test.go
)

func Test_JSON_SPBF(test *testing.T) {
	fmt.Println("Bellman-Ford Shortest Path on testgraph12:")
	g12 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.012")
	fmt.Println(spbf.SPBF(g12, "S", "T"))
	// S(=0) → A(=7) → C(=4) → B(=2) → T(=-2) true
	fmt.Println(g12.ShowPrev("T"), "/ T's StampD:", g12.FindVertexByID("T").StampD)
	fmt.Println(g12.ShowPrev("B"), "/ B's StampD:", g12.FindVertexByID("B").StampD)
	fmt.Println(g12.ShowPrev("C"), "/ C's StampD:", g12.FindVertexByID("C").StampD)
	fmt.Println(g12.ShowPrev("A"), "/ A's StampD:", g12.FindVertexByID("A").StampD)
	fmt.Println(g12.ShowPrev("S"), "/ S's StampD:", g12.FindVertexByID("S").StampD)
	/*
	   Prev of T:  A B / T's StampD: -2
	   Prev of B:  S C / B's StampD: 2
	   Prev of C:  A B T / C's StampD: 4
	   Prev of A:  S B / A's StampD: 7
	   Prev of S:  T / S's StampD: 0

	   Add the smallest from back
	   	that is not source, destination
	   	that is not in the result
	*/

	fmt.Println("Bellman-Ford Shortest Path on testgraph13:")
	g13 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.013")
	fmt.Println(spbf.SPBF(g13, "S", "T"))
	// There is negative weighted cycle (No Shortest Path) false
}
