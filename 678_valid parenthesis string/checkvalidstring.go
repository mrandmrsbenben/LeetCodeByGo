package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "((*)"
	// s := "**((*"
	// s := "(*(*))((*"
	// s := "(())((())()()(*)(*()(())())())()()((()())((()))(*"
	// s := "(*()"
	// s := "***(*(***(*****"
	s := "(()(())()())*((()(())))*()(*)()()(*((()((*(*))))()*()(()((()(*((()))*(((())(())))*))(()*))(()*)"
	// s := ")("
	fmt.Println("Output: ", checkValidString(s))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func checkValidString(s string) bool {
	for strings.Contains(s, "()") {
		s = strings.Replace(s, "()", "", -1)
	}
	max, min := 0, 0
	for i := range s {
		switch s[i] {
		case '(':
			min++
			max++
		case ')':
			if min > 0 {
				min--
			}
			max--
		case '*':
			if min > 0 {
				min--
			}
			max++
		}
		if max < 0 {
			return false
		}
	}
	return min == 0
}
