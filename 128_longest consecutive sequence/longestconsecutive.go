package main

import (
	"fmt"
)

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	// nums := []int{2147483646, -2147483647, 0, 2, 2147483644, -2147483645, 2147483645}
	fmt.Println("Output: ", longestConsecutive(nums))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了97.24%的用户
//内存消耗 :3.7 MB, 在所有 Go 提交中击败了50.00%的用户
func longestConsecutive(nums []int) int {
	nm := make(map[int]int, len(nums))
	for _, n := range nums {
		nm[n] = 1
	}
	res := 0
	var curNum, cnt int
	for _, n := range nums {
		if nm[n-1] == 0 {
			curNum = n
			cnt = 1
			for nm[curNum+1] == 1 {
				curNum++
				cnt++
			}
			if cnt > res {
				res = cnt
			}
		}
	}
	return res
}
