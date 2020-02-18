package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// nums := []int{12, 345, 2, 6, 7896}
	nums := []int{555, 901, 482, 1771, 1000, 10, 100}
	fmt.Println("Output: ", findNumbers(nums))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了93.87%的用户
//内存消耗 :3.1 MB, 在所有 Go 提交中击败了80.51%的用户
func findNumbers(nums []int) int {
	sort.Ints(nums)
	cnt := 0
	base := 10
	flag := false
	for _, n := range nums {
		for n >= base {
			base *= 10
			flag = !flag
		}
		if flag {
			cnt++
		}
	}
	return cnt
}

//执行用时 :4 ms, 在所有 Go 提交中击败了93.87%的用户
//内存消耗 :3.2 MB, 在所有 Go 提交中击败了50.00%的用户
func findNumbersV1(nums []int) int {
	cnt := 0
	for _, n := range nums {
		if len(strconv.Itoa(n))%2 == 0 {
			cnt++
		}
	}
	return cnt
}
