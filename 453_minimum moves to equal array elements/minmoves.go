package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2, 4, 9}
	fmt.Printf("Output: %d\n", minMoves(nums))
}

func minMoves(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}
	sum := 0
	for i := range nums {
		sum = sum + nums[i] - min
	}
	return sum
}

func minMoves0(nums []int) int {
	min := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		min = min + nums[i] - nums[0]
	}
	return min
}
