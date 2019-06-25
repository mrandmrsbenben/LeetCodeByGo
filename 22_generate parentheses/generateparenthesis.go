package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Output:", generateParenthesis(3))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了93.13%的用户
//内存消耗 :8.6 MB, 在所有 Go 提交中击败了15.60%的用户
func generateParenthesis(n int) []string {
	strs := make([]string, 0)
	var generate func(str string, l, r int)
	generate = func(str string, l, r int) {
		if l == 0 {
			strs = append(strs, str+strings.Repeat(")", r))
			return
		}
		generate(str+"(", l-1, r)
		if l < r {
			generate(str+")", l, r-1)
		}
	}
	generate("", n, n)
	return strs
}
