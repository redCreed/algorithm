package greedy

import (
	"testing"
)

func TestNewFGraph(t *testing.T) {
	vertex := []string{"A", "B", "C", "D", "E", "F", "G"}
	weight := [][]int{
		{0, 5, 7, 10000, 10000, 10000, 2},
		{5, 0, 10000, 9, 10000, 10000, 3},
		{7, 10000, 0, 10000, 8, 10000, 10000},
		{10000, 9, 10000, 0, 10000, 4, 10000},
		{10000, 10000, 8, 10000, 0, 5, 4},
		{10000, 10000, 10000, 4, 5, 0, 6},
		{2, 3, 10000, 10000, 4, 6, 0},
	}
	l := 7
	fg := NewFGraph(l, weight, vertex)
	fg.Floyd()
	fg.print()
}
