package main

import "fmt"

func main() {
	// common.MakeTestCases("validMountainArray")
	input1 := []int{2, 1}
	fmt.Printf("Expect1: false\n")
	fmt.Printf("Output1: %v\n", validMountainArray(input1))
	input2 := []int{3, 5, 5}
	fmt.Printf("Expect2: false\n")
	fmt.Printf("Output2: %v\n", validMountainArray(input2))
	input3 := []int{0, 3, 2, 1}
	fmt.Printf("Expect3: true\n")
	fmt.Printf("Output3: %v\n", validMountainArray(input3))
}

//执行用时 :52 ms, 在所有Go提交中击败了34.38%的用户
//内存消耗 :6.4 MB, 在所有Go提交中击败了50.00%的用户
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	} else if A[0] >= A[1] {
		return false
	}
	passPeak := false
	for i := 1; i < len(A); i++ {
		if passPeak && A[i] >= A[i-1] {
			return false
		} else if !passPeak {
			if A[i] == A[i-1] {
				return false
			} else if A[i] < A[i-1] {
				passPeak = true
			}
		}
	}
	return passPeak
}
