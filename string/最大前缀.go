package string

/*
	编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
	示例1:
	输入: ["flower","flow","flight"] 输出: "fl"
	示例 2:
	输入: ["dog","racecar","car"] 输出: "" 解释: 输入不存在公共前缀。
	解释:
	输入不存在公共前缀。
	说明：
	所有输入只包含小写字母 a-z
*/

func GetPrefix(arr []string) string {
	if len(arr) <= 1 {
		return ""
	}
	firstStr := arr[0]
	l := len(arr)
	//最长前缀，如果中间不相等直接返回就行了
	for i := range firstStr {
		for j := 1; j < l; j++ {
			if arr[j][i] != firstStr[i] {
				return firstStr[:i]
			}

		}
	}
	return ""
}
