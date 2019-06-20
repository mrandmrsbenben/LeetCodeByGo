package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Printf("Output: %v\n", intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	dic := make(map[int]int)
	s := make([]int, 0)
	for _, v := range nums2 {
		dic[v] = v
	}
	for _, v := range nums1 {
		if _, ok := dic[v]; ok {
			s = append(s, v)
			delete(dic, v)
		}
	}
	return s
}
