package main

import (
	"fmt"
	"strings"
)

func main() {
	// common.MakeTestCases("repeatedSubstringPattern")
	input1 := "abab"
	fmt.Printf("Expect1: True\n")
	fmt.Printf("Output1: %v\n", repeatedSubstringPattern(input1))
	input2 := "aabaaba"
	fmt.Printf("Expect2: False\n")
	fmt.Printf("Output2: %v\n", repeatedSubstringPattern(input2))
	input3 := "abcabcabcabc"
	fmt.Printf("Expect3: True\n")
	fmt.Printf("Output3: %v\n", repeatedSubstringPattern(input3))
}

// 执行用时 : 8 ms, 在Repeated Substring Pattern的Go提交中击败了100.00% 的用户
// 内存消耗 : 5.1 MB, 在Repeated Substring Pattern的Go提交中击败了86.67% 的用户
func repeatedSubstringPattern(s string) bool {
	var pattern bool
	for i := 0; i < len(s)/2+1; i++ {
		if i > 0 && len(s)%i != 0 {
			continue
		}
		pattern = true
		for j := i; j < len(s); j = j + i {
			if j+i > len(s) || s[0:i] != s[j:j+i] || i == 0 {
				pattern = false
				break
			}
		}
		if pattern {
			return true
		}
	}
	return false
}

// 执行用时 : 592 ms, 在Repeated Substring Pattern的Go提交中击败了27.03% 的用户
// 内存消耗 : 5.1 MB, 在Repeated Substring Pattern的Go提交中击败了86.67% 的用户
func repeatedSubstringPatternV1(s string) bool {
	if len(s) < 2 {
		return false
	}
	var cnt int
	for i := 0; i < len(s)/2+1; i++ {
		cnt = strings.Count(s, s[0:i])
		if cnt > 1 && cnt*len(s[0:i]) == len(s) {
			return true
		}
	}
	return false
}
