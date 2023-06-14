package dp

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ret := Solution(height)

	fmt.Println("ret:", ret)
}
