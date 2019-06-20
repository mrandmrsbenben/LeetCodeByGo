package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "loveleetcode"
	fmt.Printf("Output: %d\n", firstUniqChar(s))
}

// 执行用时 : 16 ms, 在First Unique Character in a String的Go提交中击败了91.79% 的用户
// 内存消耗 : 5.8 MB, 在First Unique Character in a String的Go提交中击败了34.50% 的用户
func firstUniqChar(s string) int {
	for i := 0; i < len(s); i++ {
		if strings.Index(s, string(s[i])) == i && strings.LastIndex(s, string(s[i])) == i {
			return i
		}
	}
	return -1
}

// 执行用时 : 24 ms, 在First Unique Character in a String的Go提交中击败了69.91% 的用户
// 内存消耗 : 5.8 MB, 在First Unique Character in a String的Go提交中击败了36.26% 的用户
func firstUniqCharV2(s string) int {
	var found bool
	for i := 0; i < len(s); i++ {
		found = false
		for j := 0; j < i; j++ {
			if s[i] == s[j] {
				found = true
				break
			}
		}
		if !found {
			for j := i + 1; j < len(s); j++ {
				if s[i] == s[j] {
					found = true
					break
				}
			}
			if !found {
				return i
			}
		}
	}
	return -1
}

// 执行用时 : 204 ms, 在First Unique Character in a String的Go提交中击败了5.47% 的用户
// 内存消耗 : 5.8 MB, 在First Unique Character in a String的Go提交中击败了44.44% 的用户
func firstUniqCharV1(s string) int {
	for i := 0; i < len(s); i++ {
		if strings.Count(s, s[i:i+1]) == 1 {
			return i
		}
	}
	return -1
}
