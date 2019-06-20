package main

import (
	"fmt"
	"sort"
)

func main() {
	// common.MakeTestCases("findPairs")
	// nums := []int{3, 1, 4, 1, 5}
	// k := 2
	// fmt.Printf("Expect1: 2\n")
	// fmt.Printf("Output1: %v\n", findPairs(nums, k))
	// nums := []int{1, 2, 3, 4, 5}
	// k := 1
	// fmt.Printf("Expect2: 4\n")
	// fmt.Printf("Output2: %v\n", findPairs(nums, k))
	// nums := []int{1, 3, 1, 5, 4}
	// k := 0
	// fmt.Printf("Expect3: 1\n")
	// fmt.Printf("Output3: %v\n", findPairs(nums, k))
	nums := []int{1, 1, 1, 1, 1}
	k := 0
	fmt.Printf("Output: %v\n", findPairs(nums, k))

}

//执行用时 :44 ms, 在所有Go提交中击败了39.13%的用户
//内存消耗 :6 MB, 在所有Go提交中击败了100.00%的用户
func findPairs(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	cnt := 0
	sort.Ints(nums)
	if k == 0 {
		lastVal := nums[0]
		for i := 0; i < len(nums); i++ {
			if i+1 < len(nums) && nums[i+1] == nums[i] {
				if i == 0 || nums[i] != lastVal {
					cnt++
					i++
					lastVal = nums[i]
				}
			}
		}
	} else {
		for i := 0; i < len(nums); i++ {
			if i+1 < len(nums) && nums[i+1] == nums[i] {
				continue
			}
			for j := i + 1; j < len(nums); j++ {
				if j+1 < len(nums) && nums[j+1] == nums[j] {
					continue
				}
				if nums[j]-nums[i] == k {
					cnt++
				} else if nums[j]-nums[i] > k {
					break
				}
			}
		}
	}
	return cnt
}
