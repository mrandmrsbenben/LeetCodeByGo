package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15, 24, 25},
		{2, 5, 8, 12, 19, 26, 29},
		{3, 6, 9, 16, 22, 27, 33},
		{10, 13, 14, 17, 28, 32, 35},
		{18, 23, 26, 30, 40, 42, 45}}
	target := 34
	// matrix := [][]int{{5}, {6}}
	// target := 6
	fmt.Println("Output: ", searchMatrix(matrix, target))
}

//执行用时 :32 ms, 在所有 Go 提交中击败了66.67%的用户
//内存消耗 :6.3 MB, 在所有 Go 提交中击败了75.00%的用户
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 ||
		target < matrix[0][0] ||
		target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	ret := false
	toC := len(matrix[0]) - 1
	c := 0
	for r := 0; r < len(matrix) && !ret; r++ {
		if matrix[r][0] > target {
			break
		} else if matrix[r][len(matrix[r])-1] >= target {
			c = 0
			for toC-c >= 0 {
				if matrix[r][c+(toC-c)/2] == target {
					ret = true
					break
				} else if matrix[r][c+(toC-c)/2] > target {
					toC -= (toC - c) / 2
				} else {
					c += (toC - c) / 2
				}
				if toC-c == 1 {
					if matrix[r][c] == target || matrix[r][toC] == target {
						ret = true
					}
					break
				}
			}
		}
	}
	return ret
}
