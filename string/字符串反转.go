package string

import "fmt"

func StrReverse(str string) {
	s := []rune(str)
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	fmt.Println(string(s))
}
