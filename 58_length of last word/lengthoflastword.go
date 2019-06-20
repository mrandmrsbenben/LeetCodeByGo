package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	fmt.Println("Output:", lengthOfLastWord(s))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.3 MB, 在所有 Go 提交中击败了6.52%的用户
func lengthOfLastWord(s string) int {
	if len(strings.TrimSpace(s)) == 0 {
		return 0
	}
	strs := strings.Fields(s)
	return len(strs[len(strs)-1])
}
