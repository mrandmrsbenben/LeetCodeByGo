package main

import "fmt"

func main() {
	m, n := 7, 3
	fmt.Println("Output: ", uniquePaths(m, n))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if i == m-1 || j == n-1 {
				dp[i][j] = 1
			}
		}
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}

	return dp[0][0]
}
