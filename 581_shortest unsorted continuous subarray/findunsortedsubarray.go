package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{2, 6, 4, 8, 10, 9, 15}
	// nums := []int{1, 3, 2, 2, 2}
	// nums := []int{2, 1, 3, 4, 5}
	// nums := []int{1, 2, 4, 5, 3}
	nums := []int{1, 2}
	fmt.Println("Output:", findUnsortedSubarray(nums))
}

//执行用时 :56 ms, 在所有Go提交中击败了58.33%的用户
//内存消耗 :6.3 MB, 在所有Go提交中击败了44.44%的用户
func findUnsortedSubarray(nums []int) int {
	length := len(nums)
	buf := make([]int, length)
	copy(buf, nums)
	sort.Ints(buf)
	start, end := 0, length-1
	fs, fe := false, false
	for i := 0; i < length; i++ {
		if !fs && nums[i] == buf[i] {
			start++
		} else {
			fs = true
		}
		if !fe && nums[length-1-i] == buf[length-1-i] {
			end--
		} else {
			fe = true
		}
		if fs && fe {
			break
		}
	}
	if end <= start {
		return 0
	}
	return end - start + 1
}

//执行用时 :64 ms, 在所有Go提交中击败了58.33%的用户
//内存消耗 :6.3 MB, 在所有Go提交中击败了61.11%的用户
func findUnsortedSubarrayV1(nums []int) int {
	length := len(nums)
	buf := make([]int, length)
	copy(buf, nums)
	sort.Ints(buf)
	start, end := 0, 0
	for i := 0; i < length; i++ {
		if nums[i] != buf[i] {
			start = i
			break
		}
	}
	for i := length - 1; i >= start; i-- {
		if nums[i] != buf[i] {
			end = i
			break
		}
	}
	if start == end {
		return 0
	}
	return end - start + 1
}
