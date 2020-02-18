package main

import "fmt"

func main() {
	s := "abb"
	fmt.Println("Output: ", removePalindromeSub(s))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func removePalindromeSub(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return 2
		}
	}
	return 1
}
