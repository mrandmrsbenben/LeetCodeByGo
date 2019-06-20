package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{9, 1, 5, 6, 7, 2}
	fmt.Println(arrayPairSum(nums))
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums)/2; i++ {
		sum = sum + nums[i*2]
	}
	return sum
}
