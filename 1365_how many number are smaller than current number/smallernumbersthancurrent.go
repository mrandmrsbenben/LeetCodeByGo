package main

import (
	"fmt"
	"sort"
)

func main() {

	// nums1 := []int{8, 1, 2, 2, 3}
	nums := []int{6, 5, 4, 8}
	fmt.Printf("Expect1: {4,0,1,1,3}\n")
	fmt.Printf("Output1: %v\n", smallerNumbersThanCurrent(nums))
}

func smallerNumbersThanCurrent(nums []int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	for i := range nums {
		nums[i] = sort.SearchInts(sorted, nums[i])
	}
	return nums
}
