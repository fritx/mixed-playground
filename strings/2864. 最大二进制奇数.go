package strings

import "strings"

// 方法一：贪心
func maximumOddBinaryNumber(s string) string {
	// fix: panic: strings: negative Repeat count
	// 题目OJ能通过 但本地测试用例不通过
	if s == "" {
		return ""
	}
	cnt := strings.Count(s, "1")
	return strings.Repeat("1", cnt-1) + strings.Repeat("0", len(s)-cnt) + "1"
}

// 两次遍历
func maximumOddBinaryNumber2(s string) string {
	n := len(s)
	ans := make([]rune, n)
	oneCnt := 0
	for _, c := range s {
		if c == '1' {
			oneCnt += 1
		}
	}
	for i := 0; i < n-1; i++ {
		if i < oneCnt-1 {
			ans[i] = '1'
		} else {
			ans[i] = '0'
		}
	}
	// bugfix.1 - panic: runtime error: index out of range [-1]
	// 题目OJ能通过 但本地测试用例不通过
	if n > 0 {
		ans[n-1] = '1'
	}
	return string(ans)
}
