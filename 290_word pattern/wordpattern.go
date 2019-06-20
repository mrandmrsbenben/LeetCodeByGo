package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "abba"
	// pattern := "aaaa"
	// str := "dog cat cat dog"
	// str := "dog cat cat fish"
	str := "dog dog dog dog"
	fmt.Printf("Output: %v\n", wordPattern(pattern, str))
}

// 执行用时 : 0 ms, 在Word Pattern的Go提交中击败了100.00% 的用户
// 内存消耗 : 2 MB, 在Word Pattern的Go提交中击败了41.94% 的用户
func wordPattern(pattern string, str string) bool {
	mp := make(map[string]string)
	ms := make(map[string]string)
	ap := strings.Split(pattern, "")
	as := strings.Fields(str)
	if len(ap) != len(as) {
		return false
	}
	for i := range ap {
		if v, ok := mp[ap[i]]; ok {
			if v != as[i] {
				return false
			}
		} else {
			mp[ap[i]] = as[i]
		}
		if v, ok := ms[as[i]]; ok {
			if v != ap[i] {
				return false
			}
		} else {
			ms[as[i]] = ap[i]
		}
	}
	return true
}
