package main

import (
	"fmt"
	"sort"
)

func main() {
	// arr := []int{1, 2, 2, 1, 1, 3}
	// arr := []int{1, 2}
	arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println("Output: ", uniqueOccurrences(arr))
}

//执行用时 :0 ms, 在所有 golang 提交中击败了100.00%的用户
//内存消耗 :2.2 MB, 在所有 golang 提交中击败了100.00%的用户
func uniqueOccurrences(arr []int) bool {
	sort.Ints(arr)
	om := make(map[int]int)
	cnt := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			cnt++
		} else {
			if om[cnt] == 1 {
				return false
			}
			om[cnt] = 1
			cnt = 1
		}
	}
	if om[cnt] == 1 {
		return false
	}
	return true
}
