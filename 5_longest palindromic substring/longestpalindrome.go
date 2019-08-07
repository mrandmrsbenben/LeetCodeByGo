package main

import "fmt"

func main() {
	fmt.Println("Output:", longestPalindrome("babad"))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了94.35%的用户
//内存消耗 :2.3 MB, 在所有 Go 提交中击败了64.18%的用户
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	start, end := 0, 0
	var l1, l2, l int
	for i := range s {
		l1 = expand(s, i, i)
		l2 = expand(s, i, i+1)
		if l1 > l2 {
			l = l1
		} else {
			l = l2
		}
		if l > end-start {
			start = i - (l-1)/2
			end = i + l/2
		}
	}
	return s[start : end+1]
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
