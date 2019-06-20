package main

import "fmt"

func main() {
	A := []int{0, 1, 2, 4}
	// A := []int{}
	fmt.Printf("mountain index:%d\n", peakIndexInMountainArray(A))
}

func peakIndexInMountainArray(A []int) int {
	if len(A) == 0 {
		return -1
	} else if len(A) == 1 {
		return 0
	}
	var index int
	for i := 1; i < len(A); i++ {
		if A[i] < A[i-1] {
			break
		}
		index = i
	}
	return index
}
