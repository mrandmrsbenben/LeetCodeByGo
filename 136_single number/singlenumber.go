package main

import "fmt"

func main() {
	// common.MakeTestCases("singleNumber")
	input1 := []int{2, 2, 4}
	fmt.Printf("Output: %v\n", singleNumber(input1))
	fmt.Printf("Expect: 4\n")
	input2 := []int{2, 1, 2, 1, 9}
	fmt.Printf("Output: %v\n", singleNumber(input2))
	fmt.Printf("Expect: 9\n")
}

func singleNumber(nums []int) int {
	var n int
	for _, v := range nums {
		n = n ^ v
	}
	return n
}
