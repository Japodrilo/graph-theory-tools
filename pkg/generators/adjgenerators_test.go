package generators

import (
	"github.com/Japodrilo/graph-theory-tools/pkg/graph"
	"math/rand"
	"testing"
)

func TestIsCompleteMatrixGraph(t *testing.T) {
	a := [][]byte{
		{0, 1, 1, 1, 1, 1},
		{1, 0, 1, 1, 1, 1},
		{1, 1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 0},
	}
	b := [][]byte{
		{0, 1, 1},
		{1, 1, 1},
		{1, 1, 0},
	}
	c := [][]byte{
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 0},
	}
	k := graph.NewMatrixGraph(a)
	l := graph.NewMatrixGraph(b)
	m := graph.NewMatrixGraph(c)
	if !IsCompleteMatrixGraph(k) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCompleteMatrixGraph(k),
		)
	}
	if IsCompleteMatrixGraph(l) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCompleteMatrixGraph(k),
		)
	}
	if IsCompleteMatrixGraph(m) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCompleteMatrixGraph(k),
		)
	}
}

// TestCompleteMatrixGraph calls CompleteMatrixGraph with five different
// randomly generated graphs, and checks each of them to be a complete graph by
// exploring their adjacency matrices.
func TestCompleteMatrixGraph(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Intn(1000)
		k := CompleteMatrixGraph(n)
		a := k.Adjacency()
		for i := range a {
			for j := range a[i] {
				if i == j && a[i][j] != 0 {
					t.Error("Graph is not irreflexive")
				} else if i != j && a[i][j] != 1 {
					t.Errorf("No adjacency between %v and %v", i, j)
				}
			}
		}
	}
}

// TestIsCycleMatrixGraph calls IsCycleMatrixGraph with different hardcoded
// graphs, including cycles of different lengths, disconnected 2-regular graphs,
// and other non-cycle graphs.
func TestIsCycleMatrixGraph(t *testing.T) {
	a := [][]byte{
		{0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0},
	}
	b := [][]byte{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}
	c := [][]byte{
		{0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
	}
	d := [][]byte{
		{0, 1, 0, 1, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 0, 1, 0},
	}
	e := [][]byte{
		{0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0},
	}
	g := graph.NewMatrixGraph(a)
	h := graph.NewMatrixGraph(b)
	i := graph.NewMatrixGraph(c)
	j := graph.NewMatrixGraph(d)
	k := graph.NewMatrixGraph(e)
	if !IsCycleMatrixGraph(g) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCycleMatrixGraph(g),
		)
	}
	if !IsCycleMatrixGraph(h) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCycleMatrixGraph(h),
		)
	}
	if !IsCycleMatrixGraph(i) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCycleMatrixGraph(i),
		)
	}
	if IsCycleMatrixGraph(j) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCycleMatrixGraph(j),
		)
	}
	if IsCycleMatrixGraph(k) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCycleMatrixGraph(k),
		)
	}
}

// TestCycleMatrixGraph calls CycleMatrixGraph with five different
// randomly generated numbers, and checks each of them to be a cycle by
// exploring their adjacency matrices.
func TestCycleMatrixGraph(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Intn(1000)
		if n > 2 {
			c := CycleMatrixGraph(n)
			d := c.DegreeSequence()
			for _, v := range d {
				if v != 2 {
					t.Error("The graph is not 2-regular")
				}
			}
			a := c.Adjacency()
			if a[0][1] != 1 || a[0][n-1] != 1 {
				t.Errorf(
					"Adjacencies of vertex %v are not as expected",
					0,
				)
			}
			if a[n-1][0] != 1 || a[n-1][n-2] != 1 {
				t.Errorf(
					"Adjacencies of vertex %v are not as expected",
					n-1,
				)
			}
			for i := 1; i < n-1; i++ {
				if a[i][i-1] != 1 || a[i][i+1] != 1 {
					t.Errorf(
						"Adjacencies of vertex %v are not as expected",
						i,
					)
				}
			}
		}
	}
}
