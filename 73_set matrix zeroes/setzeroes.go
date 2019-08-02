package main

import "fmt"

func main() {
	// matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println("Output:", matrix)
}

func setZeroes(matrix [][]int) {
	var r0, c0 bool
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				if i == 0 {
					r0 = true
				} else {
					matrix[i][0] = 0
				}
				if j == 0 {
					c0 = true
				} else {
					matrix[0][j] = 0
				}
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if r0 {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}
	if c0 {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}

func setZeroesV1(matrix [][]int) {
	r := make(map[int]int)
	c := make(map[int]int)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				r[i] = 0
				c[j] = 0
			}
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			if _, ok := r[i]; ok {
				matrix[i][j] = 0
			} else if _, ok := c[j]; ok {
				matrix[i][j] = 0
			}
		}
	}
}
