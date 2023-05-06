package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var a = []int{3, 4, 0, 1, 7, 5, 8, 6}
	MergeSort(a, 0, len(a)-1)
	fmt.Println(a)
}
