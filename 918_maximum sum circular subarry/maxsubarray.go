package main

import (
	"fmt"
)

func main() {
	A := []int{1, -3, 3, -2}
	// A := []int{5, -3, 5}
	fmt.Println("Output: ", maxSubarraySumCircular(A))
}

func maxSubarraySumCircular(A []int) int {
	res, cur := A[0], A[0]
	lenA := len(A)
	for i := 1; i < lenA; i++ {
		if cur > 0 {
			cur += A[i]
		} else {
			cur = A[i]
		}
		if cur > res {
			res = cur
		}
	}

	rSum := make([]int, lenA)
	rSum[lenA-1] = A[lenA-1]
	for i := lenA - 2; i >= 0; i-- {
		rSum[i] = rSum[i+1] + A[i]
	}

	maxR := make([]int, lenA)
	maxR[lenA-1] = A[lenA-1]
	for i := lenA - 2; i >= 0; i-- {
		maxR[i] = max(maxR[i+1], rSum[i])
	}

	lSum := 0
	for i := 0; i < lenA-2; i++ {
		lSum += A[i]
		res = max(res, lSum+maxR[i+2])
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
