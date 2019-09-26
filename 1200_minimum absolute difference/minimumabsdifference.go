package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Output: ", minimumAbsDifference([]int{4, 2, 1, 3}))
	// fmt.Println("Output: ", minimumAbsDifference([]int{1, 3, 6, 10, 15}))
	// fmt.Println("Output: ", minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
	fmt.Println("Output: ", minimumAbsDifference([]int{40, 11, 26, 27, -20}))
}

func minimumAbsDifference(arr []int) [][]int {
	diffs := make([][]int, 1)
	sort.Ints(arr)
	diffs[0] = []int{arr[0], arr[1]}
	min := arr[1] - arr[0]
	cnt := 1
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < min {
			min = arr[i+1] - arr[i]
			diffs = diffs[0:1]
			diffs[0] = []int{arr[i], arr[i+1]}
			cnt = 1
		} else if arr[i+1]-arr[i] == min {
			if len(diffs) > cnt {
				diffs[cnt] = []int{arr[i], arr[i+1]}
			} else {
				diffs = append(diffs, []int{arr[i], arr[i+1]})
			}
			cnt++
		}
	}
	return diffs[0:cnt]
}
