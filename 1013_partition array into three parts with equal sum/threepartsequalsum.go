package main

import "fmt"

// https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/
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
	input4 := []int{10, -10, 10, -10, 10, -10, 10, -10}
	fmt.Println("Output4:", canThreePartsEqualSum(input4))
	input5 := []int{14, 6, -10, 2, 18, -7, -4, 11}
	fmt.Println("Output5:", canThreePartsEqualSum(input5))
	input6 := []int{10, -10, 10, -10}
	fmt.Println("Output4:", canThreePartsEqualSum(input6))

}

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, n := range A {
		sum += n
	}
	if sum%3 != 0 {
		return false
	}
	sumL, sumR := A[0], A[len(A)-1]
	for i, j := 1, len(A)-2; i <= j; {
		if sumL == sumR && sumL == sum/3 {
			return true
		}
		if sumL != sum/3 {
			sumL += A[i]
			i++
		}
		if sumR != sum/3 {
			sumR += A[j]
			j--
		}
	}
	return false
}

//执行用时 :56 ms, 在所有 Go 提交中击败了61.29%的用户
//内存消耗 :6.9 MB, 在所有 Go 提交中击败了25.00%的用户
func canThreePartsEqualSumV1(A []int) bool {
	sum := 0
	for _, n := range A {
		sum += n
	}
	if sum%3 != 0 {
		return false
	}
	third := sum / 3
	sum = 0
	sum1, sum2 := false, false
	for i := range A {
		sum += A[i]
		if sum == third && !sum1 {
			sum1 = true
			sum = 0
		} else if sum == third && !sum2 {
			if i == len(A)-1 {
				return false
			}
			return true
		}
	}
	return false
}
