package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
func main() {
	nums := []int{3, 30, 34, 5, 9}
	fmt.Println("Output: ", minNumber(nums))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了85.61%的用户
//内存消耗 :2.3 MB, 在所有 Go 提交中击败了100.00%的用户
func minNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] < strs[j]+strs[i]
	})
	return strings.Join(strs, "")
}
