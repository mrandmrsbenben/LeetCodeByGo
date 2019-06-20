package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2, 3, 3}
	fmt.Printf("Output: %v\n", findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {
	dspr := make([]int, 0)
	sort.Ints(nums)
	for i := len(nums); i > 0; i-- {
		if len(nums) > 0 && nums[len(nums)-1] == i {
			nums = nums[0 : len(nums)-1]
			if len(nums) > 0 && nums[len(nums)-1] == i {
				nums = nums[0 : len(nums)-1]
			}
		} else {
			dspr = append([]int{i}, dspr...)
		}
	}
	return dspr
}
