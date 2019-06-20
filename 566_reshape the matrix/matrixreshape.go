package main

import "fmt"

func main() {
	nums := [][]int{{1, 2, 5}, {3, 4, 6}}
	r := 3
	c := 2
	fmt.Printf("Output: %v\n", matrixReshape(nums, r, c))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums)*len(nums[0]) != r*c {
		return nums
	}
	if len(nums) == r && len(nums[0]) == c {
		return nums
	}
	matrix := make([][]int, r)
	rcnt, ccnt := 0, 0
	for _, x := range nums {
		for _, y := range x {
			matrix[rcnt] = append(matrix[rcnt], y)
			ccnt = ccnt + 1
			if ccnt == c {
				ccnt = 0
				rcnt = rcnt + 1
			}
		}
	}
	return matrix
}
