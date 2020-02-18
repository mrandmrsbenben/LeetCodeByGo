package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println("Output: ", findRepeatNumber(nums))
}

//执行用时 :52 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :6.8 MB, 在所有 Go 提交中击败了100.00%的用户
func findRepeatNumber(nums []int) int {
	ret := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			ret = nums[i]
			break
		}
	}
	return ret
}
