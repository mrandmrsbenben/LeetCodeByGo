package main

import "fmt"

func main() {
	A := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("Output: %v\n", transpose(A))
}

func transpose(A [][]int) [][]int {
	TA := make([][]int, len(A[0]))
	cols := len(A)
	for i := 0; i < len(TA); i++ {
		TA[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			TA[i][j] = A[j][i]
		}
	}
	return TA
}
