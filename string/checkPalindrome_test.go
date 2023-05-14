package string

import (
	"fmt"
	"testing"
)

func TestCheckPalindrome(t *testing.T) {
	s := "张张"
	fmt.Println("first:", CheckPalindrome(s))

	s2 := "张1221张"
	fmt.Println("second:", CheckPalindrome(s2))

	s3 := "张122"
	fmt.Println("three:", CheckPalindrome(s3))
}
