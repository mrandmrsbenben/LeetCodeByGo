package main

import "fmt"

func main() {
	// matrix := [][]int{{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9}}
	matrix := [][]int{{1, 2},
		{4, 5},
		{7, 8}}
	fmt.Println("Output:", findDiagonalOrder(matrix))
}

//执行用时 :468 ms, 在所有 Go 提交中击败了97.30%的用户
//内存消耗 :9.9 MB, 在所有 Go 提交中击败了50.00%的用户
func findDiagonalOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	} else if len(matrix) == 0 {
		return make([]int, 0)
	} else if len(matrix) == 1 {
		return matrix[0]
	}
	r, c := 0, 0
	lr, lc := len(matrix), len(matrix[0])
	do := make([]int, lr*lc)
	i := 0
	goUp := true
	for i < len(do) {
		do[i] = matrix[r][c]
		if goUp {
			if r != 0 && c != lc-1 {
				r--
				c++
			} else if r == 0 && c != lc-1 {
				goUp = false
				c++
			} else if r == 0 && c == lc-1 {
				goUp = false
				r++
			} else if r != 0 && c == lc-1 {
				goUp = false
				r++
			}
		} else {
			if r != lr-1 && c != 0 {
				r++
				c--
			} else if r != lr-1 && c == 0 {
				goUp = true
				r++
			} else if r == lr-1 && c == 0 {
				goUp = true
				c++
			} else if r == lr-1 && c != 0 {
				goUp = true
				c++
			}
		}
		i++
	}
	return do
}
