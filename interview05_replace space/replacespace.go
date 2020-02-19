package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "We are happy."
	fmt.Println("Output: ", replaceSpace(s))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}
