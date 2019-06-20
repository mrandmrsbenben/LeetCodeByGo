package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{2, 7, 7, 11, 15}
	target := 14
	fmt.Printf("Output: %v\n", twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	index := make([]int, 2)
	for i := range numbers {
		index[0] = i + 1
		n := sort.SearchInts(numbers[i+1:], target-numbers[i])
		if n > len(numbers[i+1:])-1 {
			continue
		}
		if numbers[n+index[0]]+numbers[i] == target {
			index[1] = n + 1 + index[0]
			break
		}
	}
	return index
}
