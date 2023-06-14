package stack

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	arr := []int{2, 1, 5, 6, 2, 3}
	//arr := []int{2, 4}
	ret := LargestRectangleArea2(arr)
	fmt.Println("ret:", ret)
}
