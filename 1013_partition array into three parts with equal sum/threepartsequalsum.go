package main

import (
	"fmt"
)

func main() {
	// common.MakeTestCases("canThreePartsEqualSum")
	input1 := []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}
	fmt.Printf("Expect1: true\n")
	fmt.Printf("Output1: %v\n", canThreePartsEqualSum(input1))
	input2 := []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}
	fmt.Printf("Expect2: false\n")
	fmt.Printf("Output2: %v\n", canThreePartsEqualSum(input2))
	input3 := []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}
	fmt.Printf("Expect3: true\n")
	fmt.Printf("Output3: %v\n", canThreePartsEqualSum(input3))
}

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, n := range A {
		sum += n
	}
	if sum%3 != 0 {
		return false
	}
	sum1, sum2 := sum/3, sum/3
	for i := range A {
		if sum1 != 0 {
			sum1 -= A[i]
		} else if sum2 != 0 {
			sum2 -= A[i]
			if sum2 == 0 {
				return true
			}
		}
	}
	return false
}
