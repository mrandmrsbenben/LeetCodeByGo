package main

import "fmt"

// https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println("Output: ", twoSum(nums, 9))
}

//执行用时 :224 ms, 在所有 Go 提交中击败了89.04%的用户
//内存消耗 :10.4 MB, 在所有 Go 提交中击败了100.00%的用户
func twoSum(nums []int, target int) []int {
	var res []int
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			res = []int{nums[i], nums[j]}
			break
		}
	}
	return res
}
