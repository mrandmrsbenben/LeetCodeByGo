package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("Output: ", majorityElement(nums))
}

//执行用时 :24 ms, 在所有 Go 提交中击败了74.87%的用户
//内存消耗 :5.9 MB, 在所有 Go 提交中击败了100.00%的用户
func majorityElement(nums []int) int {
	m := make(map[int]int)
	major := nums[0]
	for _, n := range nums {
		if _, ok := m[n]; ok {
			m[n]++
			if m[n] > len(nums)/2 {
				major = n
				break
			}
		} else {
			m[n] = 1
		}
	}
	return major
}
