package main

import (
	"fmt"
)

func main() {
	// common.MakeTestCases("pivotIndex")
	nums1 := []int{1, 7, 3, 6, 5, 6}
	fmt.Printf("Expect1: 3\n")
	fmt.Printf("Output1: %v\n", pivotIndex(nums1))
	nums2 := []int{1, 2, 3}
	fmt.Printf("Expect2: -1\n")
	fmt.Printf("Output2: %v\n", pivotIndex(nums2))
	nums3 := []int{-1, -1, -1, 0, 1, 1}
	fmt.Printf("Output3: %v\n", pivotIndex(nums3))
}

//执行用时 :28 ms, 在所有Go提交中击败了100.00%的用户
//内存消耗 :6 MB, 在所有Go提交中击败了51.28%的用户
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		return 0
	} else if len(nums) == 2 {
		if nums[1] == 0 {
			return 0
		} else if (nums[0]) == 0 {
			return 1
		}
		return -1
	}
	sumL := 0
	sumR := 0
	for i := 1; i < len(nums); i++ {
		sumR += nums[i]
	}
	if sumL == sumR {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		sumL += nums[i-1]
		sumR -= nums[i]
		if sumL == sumR {
			return i
		}
	}
	return -1
}
