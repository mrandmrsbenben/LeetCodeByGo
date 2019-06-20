package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{-8, 3, -5, -3, -5, -2, 0}
	K := 6
	fmt.Printf("Output: %d\n", largestSumAfterKNegations(A, K))
}

func largestSumAfterKNegations(A []int, K int) int {
	abs := func(i int) int {
		if i < 0 {
			return i * -1
		}
		return i
	}
	min := abs(A[0])
	sum := 0
	sort.Ints(A)
	for _, n := range A {
		if K > 0 {
			sum = sum + abs(n)
		} else {
			sum = sum + n
		}
		if abs(n) <= min {
			min = abs(n)
		}
		if n < 0 {
			K = K - 1
		}
	}
	if K > 0 && K%2 == 1 {
		sum = sum - min*2
	}
	return sum
}

func largestSumAfterKNegationsSlow(A []int, K int) int {
	sort.Ints(A)
	sum := 0
	for i, n := range A {
		if K == 0 {
			sum = sum + n
		} else {
			if n == 0 {
				K = 0
			} else if n > 0 {
				if K%2 == 0 {
					sum = sum + n
				} else {
					sum = sum - n
				}
				K = 0
			} else {
				if K != 2 {
					sum = sum - n
					K = K - 1
				} else {
					if i == len(A)-1 {
						if K%2 == 0 {
							sum = sum + n
						} else {
							sum = sum - n
						}
					} else if A[i+1] <= 0 {
						sum = sum - n
						K = K - 1
					} else if A[i+1] > 0 && A[i+1] > -1*n {
						sum = sum + n
					} else {
						sum = sum - n
						K = K - 1
					}
				}
			}
		}
	}
	return sum
}
