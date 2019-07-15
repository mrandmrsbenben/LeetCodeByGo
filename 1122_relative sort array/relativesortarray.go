package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	// arr1 := []int{}
	// arr2 := []int{}
	fmt.Println("Output:", relativeSortArray(arr1, arr2))
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Relative Sort Array.
// Memory Usage: 6 MB, less than 100.00% of Go online submissions for Relative Sort Array.
func relativeSortArray(arr1 []int, arr2 []int) []int {
	am := make(map[int]int, len(arr2))
	for _, n := range arr2 {
		am[n] = 0
	}
	arr := make([]int, 0)
	for _, n := range arr1 {
		if _, ok := am[n]; ok {
			am[n]++
		} else {
			arr = append(arr, n)
		}
	}
	sort.Ints(arr)
	arr0 := make([]int, 0)
	for _, n := range arr2 {
		for i := 0; i < am[n]; i++ {
			arr0 = append(arr0, n)
		}
	}
	arr0 = append(arr0, arr...)
	return arr0
}
