package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 2, 4, 4}
	nums := []int{3, 2, 3, 4, 6, 5}
	fmt.Printf("Output: %v\n", findErrorNums(nums))
}

//执行用时 :36 ms, 在所有Go提交中击败了100.00%的用户
//内存消耗 :6.4 MB, 在所有Go提交中击败了80.00%的用户
func findErrorNums(nums []int) []int {
	errs := make([]int, 2)
	ns := make([]int, len(nums)+1)
	for _, v := range nums {
		if ns[v] == 1 {
			errs[0] = v
		}
		ns[v] = 1
	}
	for i := 1; i < len(ns); i++ {
		if ns[i] == 0 {
			errs[1] = i
			break
		}
	}
	return errs
}
