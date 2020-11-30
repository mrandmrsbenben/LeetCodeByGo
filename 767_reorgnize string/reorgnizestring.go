package main

import (
	"fmt"
)

func main() {
	S := "vvvlo"
	fmt.Println("Output: ", reorganizeString(S))
}

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：2 MB, 在所有 Go 提交中击败了95.45%的用户
func reorganizeString(S string) string {

	l := len(S)
	chars := make([]int, 26)
	for _, s := range S {
		chars[s-'a']++
		if chars[s-'a'] > (l+1)/2 {
			return ""
		}
	}

	strs := make([]byte, l)
	evenIdx, oddIdx, halfLen := 0, 1, l/2
	for i, c := range chars {
		ch := byte('a' + i)
		for c > 0 && c <= halfLen && oddIdx < l {
			strs[oddIdx] = ch
			c--
			oddIdx += 2
		}
		for c > 0 {
			strs[evenIdx] = ch
			c--
			evenIdx += 2
		}
	}
	return string(strs)
}
