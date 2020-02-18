package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Output: ", getNoZeroIntegers(1010))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
func getNoZeroIntegers(n int) []int {
	a := 1
	for {
		if strings.Count(strconv.Itoa(n-a), "0") == 0 && strings.Count(strconv.Itoa(a), "0") == 0 {
			return []int{a, n - a}
		}
		a++
	}
}
