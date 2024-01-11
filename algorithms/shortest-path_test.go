package algorithms_test

import (
	"testing"

	"github.com/gossie/aoc-utils/algorithms"
)

func TestShortestPath(t *testing.T) {
	n1 := algorithms.NewNode()

	n2 := algorithms.NewNode()
	n1.AddUndirectedEdge(n2, 4)

	n3 := algorithms.NewNode()
	n1.AddUndirectedEdge(n3, 3)

	n4 := algorithms.NewNode()
	n3.AddUndirectedEdge(n4, 3)

	p := algorithms.ShortestPath(n1, n4)
	if p.TotalDistance() != 6 {
		t.Fatalf("total distance was %v", p.TotalDistance())
	}
}

func TestShortestPath_multiplePossibleWays(t *testing.T) {
	n1 := algorithms.NewNode()

	n2 := algorithms.NewNode()
	n1.AddUndirectedEdge(n2, 4)

	n3 := algorithms.NewNode()
	n1.AddUndirectedEdge(n3, 3)

	n4 := algorithms.NewNode()
	n2.AddUndirectedEdge(n4, 1)
	n3.AddUndirectedEdge(n4, 3)

	p := algorithms.ShortestPath(n1, n4)
	if p.TotalDistance() != 5 {
		t.Fatalf("total distance was %v", p.TotalDistance())
	}
}

func TestShortestPath_oneWayStreet(t *testing.T) {
	n1 := algorithms.NewNode()

	n2 := algorithms.NewNode()
	n1.AddUndirectedEdge(n2, 4)

	n3 := algorithms.NewNode()
	n1.AddUndirectedEdge(n3, 3)

	n4 := algorithms.NewNode()
	n4.AddDirectedEdge(n2, 1)
	n3.AddUndirectedEdge(n4, 3)

	p := algorithms.ShortestPath(n1, n4)
	if p.TotalDistance() != 6 {
		t.Fatalf("total distance was %v", p.TotalDistance())
	}
}
