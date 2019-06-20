package main

import (
	"fmt"
)

func main() {
	// common.MakeTestCases("")
	input1 := []int{0}
	fmt.Printf("Expect1: 3\n")
	fmt.Printf("Output1: %v\n", findMaxConsecutiveOnes(input1))
}

func findMaxConsecutiveOnes(nums []int) int {
	var max, cnt int
	for _, n := range nums {
		if n == 0 {
			if max < cnt {
				max = cnt
			}
			cnt = 0
		} else {
			cnt = cnt + 1
		}
	}
	if max < cnt {
		return cnt
	}
	return max
}
