package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{1, 2, 11, 15, 2}, 4))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	s := make([]int, 2)
	for i, v1 := range nums {
		v2 := target - v1
		if _, ok := m[v2]; ok && i != m[v2] {
			return []int{m[v2], i}
		}
		m[v1] = i
	}
	return s
}
