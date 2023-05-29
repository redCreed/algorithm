package string

import (
	"fmt"
	"testing"
)

func TestViolenceMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	ret := ViolenceMaxSlidingWindow(nums, 3)

	fmt.Println("ret:", ret)
}
