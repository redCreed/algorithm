package string

import (
	"fmt"
	"testing"
)

func TestCheckBracketMatching(t *testing.T) {
	testCases := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}

	for _, testCase := range testCases {
		fmt.Printf("字符串 %s 是否对称: %v\n", testCase, CheckBracketMatching(testCase))
	}
}
