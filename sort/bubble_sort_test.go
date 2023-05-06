package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	d := []int{2, 4, 1, 5, 9}
	sorted := BubbleSort(d)
	fmt.Println(sorted)
}
