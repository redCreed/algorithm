package two_devide

import (
	"fmt"
	"testing"
)

func TestTwoDevide(t *testing.T) {
	arr := []int{1, 2, 4, 5, 6, 8, 10}
	index := TwoDevide(arr, 5)
	fmt.Println(index)
}
