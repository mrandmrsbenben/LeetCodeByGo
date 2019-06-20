package main

import (
	"fmt"
	"sort"
)

func main() {
	// common.MakeTestCases("containsDuplicate")
	input1 := []int{3, 3}
	fmt.Printf("Expect1: true\n")
	fmt.Printf("Output1: %v\n", containsDuplicate(input1))
	input2 := []int{1, 2, 3, 4}
	fmt.Printf("Expect2: false\n")
	fmt.Printf("Output2: %v\n", containsDuplicate(input2))
	input3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Printf("Expect3: true\n")
	fmt.Printf("Output3: %v\n", containsDuplicate(input3))

}

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sort.Ints(nums)
	for i := 0; i <= len(nums)-2; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
