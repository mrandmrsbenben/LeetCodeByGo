package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println("Output:", isSubsequence(s, t))
}

//执行用时 :16 ms, 在所有 Go 提交中击败了95.56%的用户
//内存消耗 :7.1 MB, 在所有 Go 提交中击败了25.00%的用户
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	for i, j := 0, 0; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
		if i == len(s) {
			return true
		}
	}
	return false
}
