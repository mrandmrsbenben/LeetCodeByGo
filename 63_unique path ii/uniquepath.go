package main

import "fmt"

func main() {
	obstacleGrid := [][]int{
		{0, 0},
		{1, 0}}
	// {0, 1}}
	// {0, 0, 0},
	// {0, 1, 0},
	// {0, 0, 0}}
	// {0, 0},
	// {1, 1},
	// {0, 0}}
	fmt.Println("Output: ", uniquePathsWithObstacles(obstacleGrid))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.5 MB, 在所有 Go 提交中击败了100.00%的用户
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		if i == m-1 {
			dp[i][n-1] = 1
			for j := n - 2; j >= 0; j-- {
				if obstacleGrid[i][j] == 1 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i][j+1]
				}
			}
		}
	}

	dp[m-1][n-1] = 1
	for i := m - 2; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				if j == n-1 {
					dp[i][j] = dp[i+1][j]
				} else {
					dp[i][j] = dp[i+1][j] + dp[i][j+1]
				}
			}
		}
	}
	return dp[0][0]
}
