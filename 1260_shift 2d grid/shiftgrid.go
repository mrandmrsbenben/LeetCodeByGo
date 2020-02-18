package main

import "fmt"

func main() {
	// grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// k := 1
	grid := [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}
	k := 4
	fmt.Println("Output: ", shiftGrid(grid, k))
}

//执行用时 :24 ms, 在所有 Go 提交中击败了98.11%的用户
//内存消耗 :6.2 MB, 在所有 Go 提交中击败了76.09%的用户
func shiftGrid(grid [][]int, k int) [][]int {
	r := len(grid)
	c := len(grid[0])
	k = k % (r * c)
	if k == 0 {
		return grid
	}

	// input all value to one slice and move
	all := make([]int, 0, r*c)
	for i := 0; i < r; i++ {
		all = append(all, grid[i]...)
	}
	all = append(all[r*c-k:], all[0:r*c-k]...)

	n := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			grid[i][j] = all[n]
			n++
		}
	}

	return grid
}

//执行用时 :24 ms, 在所有 Go 提交中击败了98.11%的用户
//内存消耗 :6.3 MB, 在所有 Go 提交中击败了41.30%的用户
func shiftGrid1(grid [][]int, k int) [][]int {
	r := len(grid)
	c := len(grid[0])
	k = k % (r * c)
	for k > 0 {
		for i := r - 1; i >= 0; i-- {
			for j := c - 1; j >= 0; j-- {
				if j > 0 {
					grid[i][j], grid[i][j-1] = grid[i][j-1], grid[i][j]
				} else if i > 0 {
					grid[i][j], grid[i-1][c-1] = grid[i-1][c-1], grid[i][j]
				}
			}
		}
		k--
	}
	return grid
}
