package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
func main() {
	arr := []int{3, 2, 1}
	k := 2
	fmt.Println("Output: ", getLeastNumbers(arr, k))
}

//执行用时 :40 ms, 在所有 Go 提交中击败了46.74%的用户
//内存消耗 :6.3 MB, 在所有 Go 提交中击败了100.00%的用户
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[0:k]
}
