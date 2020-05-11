// https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	nums := []int{0, 0, 2, 2, 4}
	fmt.Println("Output: ", isStraight(nums))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
func isStraight(nums []int) bool {
	sort.Ints(nums)
	zeroCnt := 0
	for i := range nums {
		if nums[i] == 0 {
			zeroCnt++
		} else if i > 0 && nums[i-1] > 0 {
			if nums[i] == nums[i-1] {
				return false
			} else if nums[i]-nums[i-1] > 1 {
				zeroCnt -= nums[i] - nums[i-1] - 1
				if zeroCnt < 0 {
					return false
				}
			}
		}

	}
	return true
}
