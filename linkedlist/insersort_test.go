package linkedlist

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	nums := []int{2, 5, 1, 4, 9, 3}
	sorted := InsertSort(nums)
	fmt.Println(sorted)
}
