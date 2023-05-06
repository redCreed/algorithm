package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	s := []int{2, 1, 6, 4, 7, 8, 0}
	sorted := QuickSort(s)
	fmt.Println("sorted arr:", sorted)
}
