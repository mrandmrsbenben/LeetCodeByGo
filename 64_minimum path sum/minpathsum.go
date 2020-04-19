package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}
	fmt.Println("Output: ", minPathSum(grid))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了93.18%的用户
//内存消耗 :4.5 MB, 在所有 Go 提交中击败了16.67%的用户
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dp := make([][]int, len(grid))
	for i := range grid {
		dp[i] = make([]int, len(grid[i]))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[len(grid)-1][len(grid[0])-1] = grid[len(grid)-1][len(grid[0])-1]

	var fMinSum func(i, j int) int
	fMinSum = func(i, j int) int {
		if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[i])-1 {
			return math.MaxInt64
		}
		if dp[i][j] >= 0 {
			return dp[i][j]
		}
		dp[i][j] = min(fMinSum(i+1, j), fMinSum(i, j+1)) + grid[i][j]
		return dp[i][j]
	}
	return fMinSum(0, 0)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
