package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{-1, 2, 1, -4}
	// nums := []int{1, 1, 1, 1}
	// target := 10
	nums := []int{0, 2, 1, -3}
	target := 1
	fmt.Println("Output: ", threeSumClosest(nums, target))
}

//执行用时：4 ms, 在所有 Go 提交中击败了96.20%的用户
//内存消耗：2.7 MB, 在所有 Go 提交中击败了25.00%的用户
func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			if tmp > target {
				r--
			} else if tmp < target {
				l++
			} else {
				return target
			}
			if distance(tmp, target) < distance(res, target) {
				res = tmp
			}
		}
	}
	return res
}

func distance(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
