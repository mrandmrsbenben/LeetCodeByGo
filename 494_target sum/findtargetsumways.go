package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1}
	S := 3
	fmt.Println("Output:", findTargetSumWays(nums, S))
}

//执行用时 :516 ms, 在所有 Go 提交中击败了57.14%的用户
//内存消耗 :2.2 MB, 在所有 Go 提交中击败了66.67%的用户
func findTargetSumWays(nums []int, S int) int {
	cnt := 0
	if len(nums) == 1 {
		if nums[0] == S {
			cnt++
		}
		if nums[0] == -1*S {
			cnt++
		}
		return cnt
	}
	return findTargetSumWays(nums[1:], S-nums[0]) + findTargetSumWays(nums[1:], S+nums[0])
}
