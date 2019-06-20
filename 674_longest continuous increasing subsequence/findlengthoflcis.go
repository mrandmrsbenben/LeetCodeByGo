package main

import "fmt"

func main() {
	// common.MakeTestCases("findLengthOfLCIS")
	input1 := []int{1, 3, 5, 4, 7}
	fmt.Printf("Expect1: 3\n")
	fmt.Printf("Output1: %v\n", findLengthOfLCIS(input1))
	input2 := []int{2, 2, 2, 2, 2}
	fmt.Printf("Expect2: 1\n")
	fmt.Printf("Output2: %v\n", findLengthOfLCIS(input2))
}

// 执行用时 : 12 ms, 在Longest Continuous Increasing Subsequence的Go提交中击败了100.00% 的用户
// 内存消耗 : 4.6 MB, 在Longest Continuous Increasing Subsequence的Go提交中击败了8.33% 的用户
func findLengthOfLCIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	max := 0
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			if i-start > max {
				max = i - start
			}
			start = i
			if len(nums)-start <= max {
				break
			}
		} else if i == len(nums)-1 && i-start+1 > max {
			max = i - start + 1
		}
	}
	return max
}
