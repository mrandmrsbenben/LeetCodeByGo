package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{1, 3, 6}
	K := 3
	fmt.Printf("Smallest Range: %d\n", smallestRangeI(A, K))
}

func smallestRangeI(A []int, K int) int {
	sort.Ints(A)
	r := (A[len(A)-1] - K) - (A[0] + K)
	if r < 0 {
		r = 0
	}
	return r
}
