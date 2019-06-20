package main

import (
	"fmt"
	"sort"
)

func main() {
	// common.MakeTestCases("missingNumber")
	input1 := []int{3, 0, 1}
	fmt.Printf("Expect1: 2\n")
	fmt.Printf("Output1: %v\n", missingNumber(input1))
	input2 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Printf("Expect2: 8\n")
	fmt.Printf("Output2: %v\n", missingNumber(input2))
}

func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i := range nums {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
