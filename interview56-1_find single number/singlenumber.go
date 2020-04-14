package main

import "fmt"

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
func main() {
	nums := []int{4, 1, 4, 3}
	fmt.Println("Output: ", singleNumbers(nums))
}

//执行用时 :24 ms, 在所有 Go 提交中击败了58.60%的用户
//内存消耗 :6 MB, 在所有 Go 提交中击败了100.00%的用户
func singleNumbers(nums []int) []int {
	var n int
	for i := range nums {
		n ^= nums[i]
	}
	digit1 := n & (^n + 1)
	var x int
	for i := range nums {
		if nums[i]&digit1 == 0 {
			x ^= nums[i]
		}
	}
	return []int{x, n ^ x}
}
