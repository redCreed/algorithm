package string

//判断字符串是否对称  例如 "[" "]" "{" "}" "(" ")" 等
// "{()}"  "{}"  等都对称

// CheckBracketMatching 用于判断字符串中的括号是否对称
func CheckBracketMatching(s string) bool {
	// 初始化一个栈，使用切片来模拟
	stack := []rune{}
	// 定义一个映射，用于存储右括号和对应的左括号
	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 遍历字符串中的每个字符
	for _, char := range s {
		switch char {
		// 遇到左括号，将其压入栈中
		case '(', '[', '{':
			stack = append(stack, char)
		// 遇到右括号
		case ')', ']', '}':
			// 如果栈为空，说明没有对应的左括号，返回 false
			if len(stack) == 0 {
				return false
			}
			// 弹出栈顶元素
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 检查弹出的左括号是否与当前右括号匹配
			if mapping[char] != top {
				return false
			}
		}
	}
	// 遍历结束后，如果栈为空，说明括号对称，返回 true
	return len(stack) == 0
}
