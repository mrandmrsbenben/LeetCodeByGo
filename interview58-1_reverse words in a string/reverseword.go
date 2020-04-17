package main

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
func main() {
	// s := "the sky is   blue "
	s := "  hello world!  "
	fmt.Println("Output: ", reverseWords(s))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了39.37%的用户
//内存消耗 :6.8 MB, 在所有 Go 提交中击败了100.00%的用户
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return ""
	}
	words := strings.Split(s, " ")
	i, j := 0, len(words)-1
	for i < j {
		if len(words[i]) == 0 {
			i++
		} else if len(words[j]) == 0 {
			j--
		} else {
			words[i], words[j] = words[j], words[i]
			i++
			j--
		}
	}
	s = ""
	for _, w := range words {
		if len(w) > 0 {
			s += w + " "
		}
	}
	return s[0 : len(s)-1]
}
