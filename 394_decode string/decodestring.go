package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// s := "3[a]2[bc]"
	// s := "3[a2[c]]"
	s := "2[abc]3[cd]ef"
	// s := "3[a2[c4[gh]f]d5[e]]"
	fmt.Println("Output:", decodeString(s))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
func decodeString(s string) string {
	if s == "" {
		return ""
	}
	res := ""
	numStack := make([]int, 0)
	strStack := make([]string, 0)
	nstart := -1
	var n int
	var str string
	for i := range s {
		if s[i] == '[' {
			n, _ = strconv.Atoi(s[nstart:i])
			numStack = append(numStack, n)
			strStack = append(strStack, "")
			nstart = -1
		} else if s[i] == ']' {
			n = numStack[len(numStack)-1]
			str = strings.Repeat(strStack[len(strStack)-1], n)
			strStack = strStack[0 : len(strStack)-1]
			if len(strStack) == 0 {
				res += str
			} else {
				strStack[len(strStack)-1] += str
			}
			numStack = numStack[0 : len(numStack)-1]
		} else if nstart == -1 && unicode.IsNumber(rune(s[i])) {
			nstart = i
		} else if unicode.IsLetter(rune(s[i])) {
			if len(strStack) > 0 {
				strStack[len(strStack)-1] += string(s[i])
			} else {
				res += string(s[i])
			}
		}
	}
	return res
}
