package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 5
	fmt.Printf("Output: %d\n", search(nums, target))
}

func search(nums []int, target int) int {
	if target < nums[0] || target > nums[len(nums)-1] {
		return -1
	} else if len(nums) == 1 {
		return 0
	}
	mid := len(nums) / 2
	index := search(nums[0:mid], target)
	if index == -1 {
		index = search(nums[mid:], target)
		if index != -1 {
			index = index + mid
		}
	}
	return index
}

func search1(nums []int, target int) int {
	index := sort.SearchInts(nums, target)
	if index < len(nums) && nums[index] == target {
		return index
	}
	return -1
}
