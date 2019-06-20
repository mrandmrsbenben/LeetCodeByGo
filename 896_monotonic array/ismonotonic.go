package main

import (
	"fmt"
	"sort"
)

func main() {
	// common.MakeTestCases("isMonotonic")
	input1 := []int{1, 2, 2, 3}
	fmt.Printf("Expect1: true\n")
	fmt.Printf("Output1: %v\n", isMonotonic(input1))
	input2 := []int{6, 5, 4, 4}
	fmt.Printf("Expect2: true\n")
	fmt.Printf("Output2: %v\n", isMonotonic(input2))
	input3 := []int{1, 3, 2, 4}
	fmt.Printf("Expect3: false\n")
	fmt.Printf("Output3: %v\n", isMonotonic(input3))
	input4 := []int{1, 2, 4, 5}
	fmt.Printf("Expect4: true\n")
	fmt.Printf("Output4: %v\n", isMonotonic(input4))
	input5 := []int{1, 1, 1}
	fmt.Printf("Expect5: true\n")
	fmt.Printf("Output5: %v\n", isMonotonic(input5))
}

func isMonotonic(A []int) bool {
	B := make([]int, len(A))
	copy(B, A)
	sort.Ints(B)
	inc, dec := true, true
	for i := range A {
		if inc && A[i] != B[i] {
			inc = false
		}
		if dec && A[i] != B[len(B)-1-i] {
			dec = false
		}
		if !inc && !dec {
			return false
		}
	}
	return true
}
