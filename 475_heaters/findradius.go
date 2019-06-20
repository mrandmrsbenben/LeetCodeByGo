package main

import (
	"fmt"
	"sort"
)

func main() {
	// houses := []int{1, 2, 3}
	// heaters := []int{2}
	houses := []int{1, 2, 3, 5}
	heaters := []int{1}
	fmt.Println("Output:", findRadius(houses, heaters))
}

//执行用时 :88 ms, 在所有 Go 提交中击败了65.00%的用户
//内存消耗 :6.7 MB, 在所有 Go 提交中击败了60.00%的用户
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var index, min, max int
	var left, right int
	for i := range houses {
		index = sort.SearchInts(heaters, houses[i])
		if index == 0 {
			min = heaters[index] - houses[i]
		} else if index == len(heaters) {
			min = houses[i] - heaters[len(heaters)-1]
		} else {
			left = houses[i] - heaters[index-1]
			right = heaters[index] - houses[i]
			if left > right {
				min = right
			} else {
				min = left
			}
		}
		if max < min {
			max = min
		}
	}
	return max
}
