package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	fmt.Println("Output:", spiralOrder(matrix))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.2 MB, 在所有 Go 提交中击败了7.69%的用户
func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	} else if len(matrix) == 0 {
		return []int{}
	} else if len(matrix) == 1 {
		return matrix[0]
	}
	lr, lc := len(matrix), len(matrix[0])
	so := make([]int, lr*lc)
	flag := make([][]int, lr)
	for i := range flag {
		flag[i] = make([]int, lc)
	}
	i := 0
	r, c := 0, 0
	direction := 0 //0:E 1:S 2:W 3:N
	for i < lr*lc {
		so[i] = matrix[r][c]
		flag[r][c] = 1
		switch direction {
		case 0:
			if c == lc-1 || flag[r][c+1] == 1 {
				direction = 1
				r++
			} else {
				c++
			}
		case 1:
			if r == lr-1 || flag[r+1][c] == 1 {
				direction = 2
				c--
			} else {
				r++
			}
		case 2:
			if c == 0 || flag[r][c-1] == 1 {
				direction = 3
				r--
			} else {
				c--
			}
		case 3:
			if r == 0 || flag[r-1][c] == 1 {
				direction = 0
				c++
			} else {
				r--
			}
		}
		i++
	}
	return so
}
