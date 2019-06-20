package main

import (
	"fmt"
)

func main() {
	input1 := []int{3, 2, 3}
	fmt.Printf("Expect1: 3\n")
	fmt.Printf("Output1: %v\n", majorityElement(input1))
	input2 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("Expect2: 2\n")
	fmt.Printf("Output2: %v\n", majorityElement(input2))
}

func majorityElement(nums []int) int {
	var major int
	m := make(map[int]int, 0)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			m[n] = m[n] + 1
		} else {
			m[n] = 1
		}
		if m[n] > len(nums)/2 {
			major = n
			break
		}
	}
	return major
}
