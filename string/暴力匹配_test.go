package string

import (
	"fmt"
	"testing"
)

func TestViolentMatch(t *testing.T) {
	s1 := "1basddt"
	s2 := "asd"
	index := ViolentMatch(s1, s2)
	fmt.Println("index:", index)
}
