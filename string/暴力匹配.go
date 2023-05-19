package string

// ViolentMatch 暴力匹配
func ViolentMatch(s1, s2 string) int {
	arr1 := []rune(s1)
	arr2 := []rune(s2)
	i := 0 //指向s1
	j := 0 //指向s2
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			i++
			j++
		} else {
			i = i - (j - 1)
			j = 0
		}
	}

	if j == len(arr2) {
		return i - j
	} else {
		return -1
	}
}
