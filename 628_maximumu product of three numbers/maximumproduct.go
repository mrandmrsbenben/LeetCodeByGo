package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, -2, -3, -4}
	fmt.Printf("Output: %d\n", maximumProduct(nums))
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	last := len(nums) - 1
	if nums[last] == 0 {
		return 0
	} else if nums[last] < 0 {
		return nums[last] * nums[last-1] * nums[last-2]
	} else {
		if nums[0]*nums[1] >= nums[last-1]*nums[last-2] {
			return nums[0] * nums[1] * nums[last]
		}
		return nums[last] * nums[last-1] * nums[last-2]
	}
}
