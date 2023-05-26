package string

import "fmt"

func KmpNext(str string) []int {
	next := make([]int, len(str))
	next[0] = 0
	j := 0
	for i := 1; i < len(str); i++ {
		for j > 0 && str[i] != str[j] {
			j = next[j-1]
		}

		if str[i] == str[j] {
			j++
		}
		next[i] = j
	}

	//打印最大匹配表
	fmt.Println(next)
	return next
}

// KmpSearch 返回在str1中匹配str2的下标
func KmpSearch(str1, str2 string, next []int) int {
	j := 0
	for i := 0; i < len(str1); i++ {
		for j > 0 && str1[i] != str2[j] {
			j = next[j-1]
		}

		if str1[i] == str2[j] {
			j++
		}

		if j == len(str2) {
			return i - j + 1
		}
	}

	return -1
}
