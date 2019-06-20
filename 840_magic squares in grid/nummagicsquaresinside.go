package main

import "fmt"

func main() {
	grid := [][]int{{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2}}
	// grid :=
	// 	[][]int{{5, 5, 5},
	// 		{5, 5, 5},
	// 		{5, 5, 5}}
	// grid :=
	// 	[][]int{{10, 3, 5},
	// 		{1, 6, 11},
	// 		{7, 9, 2}}
	fmt.Println("Output:", numMagicSquaresInside(grid))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了84.62%的用户
//内存消耗 :2.1 MB, 在所有 Go 提交中击败了90.00%的用户
func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}
	cnt := 0
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[0])-2; j++ {
			g := [][]int{grid[i][j : j+3], grid[i+1][j : j+3], grid[i+2][j : j+3]}
			if isMagicSquare(g) {
				cnt++
			}
		}
	}
	return cnt
}

func isMagicSquare(g [][]int) bool {
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if isValid(g[0][0], nums) || isValid(g[0][1], nums) || isValid(g[0][2], nums) {
		return false
	}
	sum := g[0][0] + g[0][1] + g[0][2]
	if sum != g[0][0]+g[1][1]+g[2][2] || sum != g[0][2]+g[1][1]+g[2][0] {
		return false
	}
	for i := 1; i < 3; i++ {
		if isValid(g[i][0], nums) || isValid(g[i][1], nums) ||
			isValid(g[i][2], nums) || sum != g[i][0]+g[i][1]+g[i][2] {
			return false
		}
	}
	for j := 0; j < 3; j++ {
		if sum != g[0][j]+g[1][j]+g[2][j] {
			return false
		}
	}
	return true
}

func isValid(n int, nums []int) bool {
	if nums[n] == 1 || n < 1 || n > 9 {
		return true
	}
	nums[n] = 1
	return false
}
