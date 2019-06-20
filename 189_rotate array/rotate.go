package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	// k := 3
	nums := []int{1, 2}
	k := 3
	rotate(nums, k)
	fmt.Printf("Output: %v\n", nums)
}

// 执行用时 : 120 ms, 在Rotate Array的Go提交中击败了96.39% 的用户
// 内存消耗 : 8.1 MB, 在Rotate Array的Go提交中击败了56.32% 的用户
func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	k = k % len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[0:len(nums)-k]...))
}
