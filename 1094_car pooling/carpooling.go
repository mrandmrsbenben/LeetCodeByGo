package main

import (
	"fmt"
	"sort"
)

func main() {
	// trips := [][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}
	// capacity := 11
	trips := [][]int{{2, 1, 5}, {3, 5, 7}}
	capacity := 3
	fmt.Println("Output: ", carPooling(trips, capacity))
}

func carPooling(trips [][]int, capacity int) bool {
	sort.Slice(trips, func(i, j int) bool {
		if trips[i][1] == trips[j][1] {
			return trips[i][2] < trips[j][2]
		}
		return trips[i][1] < trips[j][1]
	})

	var on, off [1001]int
	for _, t := range trips {
		on[t[1]] += t[0]
		off[t[2]] += t[0]
	}
	begin, end := trips[0][1], trips[len(trips)-1][2]
	sum := 0
	for i := begin; i <= end; i++ {
		sum -= off[i]
		sum += on[i]
		if sum > capacity {
			return false
		}
	}
	return true
}
