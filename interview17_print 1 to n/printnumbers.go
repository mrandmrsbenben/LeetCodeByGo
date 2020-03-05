package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Output: ", printNumbers(2))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了95.07%的用户
//内存消耗 :6.1 MB, 在所有 Go 提交中击败了100.00%的用户
func printNumbers(n int) []int {
	l := 1
	for i := 0; i < n; i++ {
		l *= 10
	}
	nums := make([]int, l-1)
	for i := 1; i < l; i++ {
		nums[i-1] = i
	}
	return nums
}

//执行用时 :8 ms, 在所有 Go 提交中击败了95.07%的用户
//内存消耗 :6.1 MB, 在所有 Go 提交中击败了100.00%的用户
func printNumbersV1(n int) []int {
	l, _ := strconv.Atoi("1" + strings.Repeat("0", n))
	nums := make([]int, l-1)
	for i := 1; i < l; i++ {
		nums[i-1] = i
	}
	return nums
}
