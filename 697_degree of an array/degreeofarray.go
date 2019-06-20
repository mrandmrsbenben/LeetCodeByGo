package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 1, 4, 2}
	fmt.Printf("Output: %d\n", findShortestSubArray(nums))
}

func findShortestSubArray(nums []int) int {
	degree := 0
	md := make(map[int][]int)
	for i, n := range nums {
		if _, ok := md[n]; ok {
			md[n] = append(md[n], i)
		} else {
			md[n] = []int{i}
		}
		if len(md[n]) > degree {
			degree = len(md[n])
		}
	}
	if degree == 1 {
		return 1
	}
	max := len(nums)
	for _, v := range md {
		if len(v) != degree {
			continue
		}
		if v[len(v)-1]-v[0]+1 < max {
			max = v[len(v)-1] - v[0] + 1
		}
	}
	return max
}
