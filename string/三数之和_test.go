package string

import (
	"fmt"
	"testing"
)

func TestViolenceThreeSum(t *testing.T) {
	nums := []int{-1, 2, 1, 4}
	ViolenceThreeSum(nums, 4)

	//使用双指针
	v := PointerThreeSum(nums, 2)
	fmt.Println(v)
}
