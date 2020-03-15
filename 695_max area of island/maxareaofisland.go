package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	// grid := [][]int{{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}}
	fmt.Println("Output: ", maxAreaOfIsland(grid))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了94.69%的用户
//内存消耗 :5.6 MB, 在所有 Go 提交中击败了12.50%的用户
func maxAreaOfIsland(grid [][]int) int {
	max := 0
	dp := make([][]int, len(grid))
	for i := range grid {
		dp[i] = make([]int, len(grid[i]))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var findMax func(i, j int) int
	findMax = func(i, j int) int {
		if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[i])-1 || grid[i][j] == 0 {
			return 0
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		dp[i][j] = 0
		area := 1 + findMax(i-1, j) + findMax(i+1, j) + findMax(i, j-1) + findMax(i, j+1)
		if area > max {
			max = area
		}
		return area
	}

	for i := range grid {
		for j := range dp[i] {
			if dp[i][j] == -1 {
				findMax(i, j)
			}
		}
	}
	return max
}
