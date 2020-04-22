package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// BinaryMatrix binary matrix
type BinaryMatrix = common.BinaryMatrix

func main() {
	// mat := [][]int{{0, 0}, {1, 1}}
	mat := [][]int{{0, 0, 0, 1}, {0, 0, 1, 1}, {0, 1, 1, 1}}
	fmt.Println("Output: ", leftMostColumnWithOne(common.CreateBinaryMatrix(mat)))
}

/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */
func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dims := binaryMatrix.Dimensions()
	lm := dims[1]
	oneFlag := false
	for i := dims[0] - 1; i >= 0; i-- {
		for j := lm - 1; j >= 0; j-- {
			if binaryMatrix.Get(i, j) == 1 {
				lm = j
				oneFlag = true
			} else {
				break
			}
		}
		if lm == 0 {
			break
		}
	}
	if oneFlag {
		return lm
	}
	return -1
}

func leftMostColumnWithOneV0(binaryMatrix BinaryMatrix) int {
	dims := binaryMatrix.Dimensions()
	lm := dims[1]
	oneFlag := false
	for i := 0; i < dims[0]; i++ {
		for j := 0; j < lm; j++ {
			if binaryMatrix.Get(i, j) == 1 {
				lm = j
				oneFlag = true
				break
			}
		}
		if lm == 0 {
			break
		}
	}
	if oneFlag {
		return lm
	}
	return -1
}
