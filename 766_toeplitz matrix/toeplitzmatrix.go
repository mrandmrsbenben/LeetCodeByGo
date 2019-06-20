package main

import "fmt"

func main() {
	// matrix := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	matrix := [][]int{{1, 2}}
	fmt.Printf("Output: %v\n", isToeplitzMatrix(matrix))
}

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 1 {
		return true
	}
	isToep := true
	for i := len(matrix) - 2; i >= 0 && isToep; i-- {
		for j := range matrix[i] {
			if j == len(matrix[i])-1 {
				break
			}
			if matrix[i][j] != matrix[i+1][j+1] {
				isToep = false
				break
			}
		}
	}
	return isToep
}
