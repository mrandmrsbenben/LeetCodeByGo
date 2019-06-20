package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("Output: %d\n", removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == nums[i+1] {
			nums = append(nums[0:i], nums[i+1:]...)
		}
	}
	return len(nums)
}
