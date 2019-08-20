package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("   the sky  is  blue  "))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("  "))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了86.45%的用户
//内存消耗 :7.4 MB, 在所有 Go 提交中击败了38.03%的用户
func reverseWords(s string) string {
	rs := ""
	if len(s) == 0 || len(strings.TrimSpace(s)) == 0 {
		return rs
	}
	start := -1
	for i := range s {
		if s[i] == 32 {
			if start != -1 {
				rs = " " + s[start:i] + rs
				start = -1
			}
			continue
		} else if start == -1 {
			start = i
		}
	}
	if start != -1 {
		rs = " " + s[start:] + rs
	}
	return rs[1:]
}

//执行用时 :20 ms, 在所有 Go 提交中击败了23.87%的用户
//内存消耗 :7.4 MB, 在所有 Go 提交中击败了40.84%的用户
func reverseWordsV1(s string) string {
	rs := ""
	if len(s) == 0 || len(strings.TrimSpace(s)) == 0 {
		return rs
	}
	words := strings.Fields(s)
	for i := len(words) - 1; i >= 0; i-- {
		rs += " " + words[i]
	}
	return rs[1:]
}
