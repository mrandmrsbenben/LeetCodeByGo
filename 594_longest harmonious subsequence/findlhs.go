package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{1, 1}
	nums := []int{1, 4, 1, 3, 1, -14, 1, -13}
	fmt.Printf("Output: %d\n", findLHS(nums))
}

// 执行用时 : 128 ms, 在Longest Harmonious Subsequence的Go提交中击败了72.73% 的用户
// 内存消耗 : 6.6 MB, 在Longest Harmonious Subsequence的Go提交中击败了55.56% 的用户
func findLHS(nums []int) int {
	sort.Ints(nums)
	max := 0
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = i
		}
		if i > 0 && nums[i] != nums[i-1] {
			if index, ok := m[nums[i-1]-1]; ok && i-index > max {
				max = i - index
			}
		}
		if i == len(nums)-1 {
			if index, ok := m[nums[i]-1]; ok && i-index+1 > max {
				max = i - index + 1
			}
		}
	}
	return max
}

// 执行用时 : 112 ms, 在Longest Harmonious Subsequence的Go提交中击败了81.82% 的用户
// 内存消耗 : 6.4 MB, 在Longest Harmonious Subsequence的Go提交中击败了88.89% 的用户
func findLHSv1(nums []int) int {
	sort.Ints(nums)
	max := 0
	start0, start1 := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[start0] == 1 {
			if nums[i] != nums[start1] {
				start1 = i
			}
			if i == len(nums)-1 && i-start0+1 > max {
				max = i - start0 + 1
			}
		} else if nums[i] == nums[start0] {
			continue
		} else {
			if i-start0 > max && nums[start0] != nums[start1] {
				max = i - start0
			}
			if nums[i]-nums[i-1] == 1 {
				start0 = start1
				start1 = i
			} else {
				start0 = i
				start1 = i
			}
		}
	}
	return max
}
