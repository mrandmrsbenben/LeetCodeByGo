package main

import "fmt"

// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Output: ", maxSubArray(nums))
}

//执行用时 :28 ms, 在所有 Go 提交中击败了66.67%的用户
//内存消耗 :6 MB, 在所有 Go 提交中击败了100.00%的用户
func maxSubArray(nums []int) int {
	sum := 0
	allsum := nums[0]
	for _, n := range nums {
		if sum < 0 {
			sum = n
		} else {
			sum = sum + n
		}
		if sum > allsum {
			allsum = sum
		}
	}
	return allsum
}
