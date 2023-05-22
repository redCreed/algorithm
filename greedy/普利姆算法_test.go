package greedy

import "testing"

func TestNewMGraph(t *testing.T) {
	minTree := &MinTree{}
	data := []string{"A", "B", "C", "D", "E", "F", "G"}
	mg := NewMGraph(len(data))
	weight := [][]int{
		{10000, 5, 7, 10000, 10000, 10000, 2},
		{5, 10000, 10000, 9, 10000, 10000, 3},
		{7, 10000, 10000, 10000, 8, 10000, 10000},
		{10000, 9, 10000, 10000, 10000, 4, 10000},
		{10000, 10000, 8, 10000, 10000, 5, 4},
		{10000, 10000, 10000, 4, 5, 10000, 6},
		{2, 3, 10000, 10000, 4, 6, 10000},
	}
	minTree.createGraph(mg, data, weight)
}
