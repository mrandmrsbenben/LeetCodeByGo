package main

import (
	"fmt"
	"sort"
)

func main() {
	// common.MakeTestCases("largestPerimeter")
	input1 := []int{2, 1, 2}
	fmt.Printf("Expect1: 5\n")
	fmt.Printf("Output1: %v\n", largestPerimeter(input1))
	input2 := []int{1, 2, 1}
	fmt.Printf("Expect2: 0\n")
	fmt.Printf("Output2: %v\n", largestPerimeter(input2))
	input3 := []int{3, 2, 3, 4}
	fmt.Printf("Expect3: 10\n")
	fmt.Printf("Output3: %v\n", largestPerimeter(input3))
	input4 := []int{3, 6, 2, 3}
	fmt.Printf("Expect4: 8\n")
	fmt.Printf("Output4: %v\n", largestPerimeter(input4))

}

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 1; i-2 >= 0; i-- {
		if A[i-1]+A[i-2] > A[i] {
			return A[i-1] + A[i-2] + A[i]
		}
	}
	return 0
}
