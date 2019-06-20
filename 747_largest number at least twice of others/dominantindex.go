package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 6, 1, 0}
	// nums := []int{4, 3, 2, 1}
	fmt.Printf("Output: %d\n", dominantIndex(nums))
}

// 执行用时 : 4 ms, 在Largest Number At Least Twice of Others的Go提交中击败了90.70% 的用户
// 内存消耗 : 2.3 MB, 在Largest Number At Least Twice of Others的Go提交中击败了51.61% 的用户
func dominantIndex(nums []int) int {
	var max, secondMax, index int
	for i, n := range nums {
		if n > max {
			secondMax, max = max, n
			index = i
		} else if n > secondMax {
			secondMax = n
		}
	}
	if max >= secondMax*2 {
		return index
	}
	return -1
}

// 执行用时 : 4 ms, 在Largest Number At Least Twice of Others的Go提交中击败了90.70% 的用户
// 内存消耗 : 2.3 MB, 在Largest Number At Least Twice of Others的Go提交中击败了22.58% 的用户
func dominantIndexV1(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	buf := make([]int, len(nums))
	copy(buf, nums)
	sort.Ints(buf)
	if buf[len(buf)-1] >= buf[len(buf)-2]*2 {
		for i := range nums {
			if nums[i] == buf[len(buf)-1] {
				return i
			}
		}
	}
	return -1
}
