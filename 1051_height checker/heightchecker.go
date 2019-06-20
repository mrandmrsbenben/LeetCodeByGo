package main

import (
	"fmt"
	"sort"
)

func main() {
	heights := []int{1, 1, 4, 2, 1, 3, 1}
	fmt.Printf("Output: %d\n", heightChecker(heights))
}

func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)
	cnt := 0
	for i := range sorted {
		if sorted[i] != heights[i] {
			cnt++
		}
	}
	return cnt
}
