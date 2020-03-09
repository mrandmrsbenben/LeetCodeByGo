package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Output: ", generateTheString(3))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.1 MB, 在所有 Go 提交中击败了100.00%的用户
func generateTheString(n int) string {
	res := ""
	if n%2 == 1 {
		res = "a"
		n--
	}
	if n > 0 {
		res += "b" + strings.Repeat("c", n-1)
	}
	return res
}
