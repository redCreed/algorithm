package string

//求字符串中的最大回文字符串 例如 "abad" 返回"aba"

// longestPalindrome 使用动态规划法求解最长回文子串
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	// 初始化二维数组 dp
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	// 单个字符都是回文串
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	start, maxLen := 0, 1
	// 枚举子串长度
	for l := 2; l <= n; l++ {
		// 枚举子串的起始位置
		for i := 0; i < n; i++ {
			j := i + l - 1
			if j >= n {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+maxLen]
}

//初始化 dp 数组：dp[i][j] 表示子串 s[i:j+1] 是否为回文串，单个字符都是回文串，所以 dp[i][i] 初始化为 true。
//状态转移方程：如果 s[i] 不等于 s[j]，则 dp[i][j] 为 false；否则，如果 j - i < 3，则 dp[i][j] 为 true，
//否则 dp[i][j] 取决于 dp[i+1][j-1]。
//记录最长回文子串：在填充 dp 数组的过程中，记录最长回文子串的起始位置和长度。

// longestPalindrome2 使用中心扩展法求解最长回文子串
func longestPalindrome2(s string) string {
	if len(s) < 1 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 以单个字符为中心扩展
		len1 := expandAroundCenter(s, i, i)
		// 以两个相邻字符为中心扩展
		len2 := expandAroundCenter(s, i, i+1)
		// 取两种情况中的最大长度
		maxLen := max(len1, len2)
		if maxLen > end-start+1 {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

// expandAroundCenter 从中心向两边扩展，返回回文串的长度
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// max 返回两个整数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//longestPalindrome2 函数：首先判断字符串长度是否小于 1，如果是则直接返回空字符串。
//然后遍历字符串的每个字符，分别以单个字符和两个相邻字符为中心调用 expandAroundCenter 函数进行扩展，
//记录最大回文串的起始和结束位置。
//expandAroundCenter 函数：从给定的中心位置向两边扩展，只要左右字符相等就继续扩展，直到不满足条件为止，
//最后返回回文串的长度。max 函数：用于比较两个整数的大小，返回较大的值。
