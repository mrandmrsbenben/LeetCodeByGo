package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{5, 4, 3}
	fmt.Printf("Output: %v\n", findRelativeRanks(nums))
}

func findRelativeRanks(nums []int) []string {
	ranks := make([]string, len(nums))
	medals := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	m := make(map[int]string)
	cp := make([]int, len(nums))
	copy(cp, nums)
	sort.Ints(cp)
	for i := len(cp) - 1; i >= 0; i-- {
		if i >= len(cp)-3 {
			m[cp[i]] = medals[len(cp)-1-i]
		} else {
			m[cp[i]] = strconv.Itoa(len(cp) - i)
		}
	}
	for i := range nums {
		ranks[i] = m[nums[i]]
	}
	return ranks
}
