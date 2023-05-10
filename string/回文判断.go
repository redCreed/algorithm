package string

/*
题目:　判断字符串是否是回文
如：abba, 从前读到尾和从尾读取前都是一样的.
abcba, aviva 都是回文字

提供三种解题思路:
1.第一种是借用一个栈, 将一半的数据存入栈中, 后半部分与出栈的元素比较.
2.第二种二个指针向中间逼近
3.第三种也是使用二个指针,先从中间开始向左右扩散
*/

// CheckPalindrome 使用第2种示例,判断是否是回文
func CheckPalindrome(s string) bool {
	if len(s) <= 1 {
		return false
	}

	//将utf8转为unicode  例如：s:="张三" utf8 len(s)=6  unicode len([]rune(s))=2
	data := []rune(s)
	left, right := 0, len(data)-1
	for left < right {
		if data[left] != data[right] {
			return false
		}
		left++
		right--
	}

	return true
}
