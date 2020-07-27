package main

import "fmt"

func main() {
	nums := [][]int{
		// {9, 9, 4},
		// {6, 6, 8},
		// {2, 1, 1}}
		// {3, 4, 5},
		// {3, 2, 6},
		// {2, 2, 1}}
		{7, 8, 9},
		{9, 7, 6},
		{7, 2, 3}}
	fmt.Println("Output: ", longestIncreasingPath(nums))
}

//执行用时：40 ms, 在所有 Go 提交中击败了73.74%的用户
//内存消耗：6.3 MB, 在所有 Go 提交中击败了80.00%的用户
/*
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	res := 0
	r, c := len(matrix), len(matrix[0])
	flag := make([][]bool, r)
	dp := make([][]int, r)
	for i := range matrix {
		flag[i] = make([]bool, c)
		dp[i] = make([]int, c)
	}

	var longestPath func(val, i, j int) int
	longestPath = func(val, i, j int) int {
		if i < 0 || i == r || j < 0 || j == c {
			return 0
		}

		if val >= matrix[i][j] {
			return 0
		}

		if flag[i][j] {
			return dp[i][j]
		}
		flag[i][j] = true

		rtn := max(max(longestPath(matrix[i][j], i-1, j), longestPath(matrix[i][j], i+1, j)),
			max(longestPath(matrix[i][j], i, j-1), longestPath(matrix[i][j], i, j+1)))
		dp[i][j] = rtn + 1
		return rtn + 1
	}

	for i := range matrix {
		for j := range matrix[i] {
			dp[i][j] = max(max(longestPath(matrix[i][j], i-1, j), longestPath(matrix[i][j], i+1, j)),
				max(longestPath(matrix[i][j], i, j-1), longestPath(matrix[i][j], i, j+1))) + 1
			if res < dp[i][j] {
				res = dp[i][j]
				// fmt.Println(res, i, j)
			}
		}
	}
	return res
}
*/
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	res := 0
	r, c := len(matrix), len(matrix[0])
	dp := make([][]int, r)
	for i := range matrix {
		dp[i] = make([]int, c)
	}

	var longestPath func(val, i, j int) int
	longestPath = func(val, i, j int) int {
		if i < 0 || i == r || j < 0 || j == c {
			return 0
		}

		if val >= matrix[i][j] {
			return 0
		}

		if dp[i][j] > 0 {
			return dp[i][j]
		}
		dp[i][j] = 1

		rtn := max(max(longestPath(matrix[i][j], i-1, j), longestPath(matrix[i][j], i+1, j)),
			max(longestPath(matrix[i][j], i, j-1), longestPath(matrix[i][j], i, j+1)))
		dp[i][j] = rtn + 1
		return rtn + 1
	}

	for i := range matrix {
		for j := range matrix[i] {
			dp[i][j] = max(max(longestPath(matrix[i][j], i-1, j), longestPath(matrix[i][j], i+1, j)),
				max(longestPath(matrix[i][j], i, j-1), longestPath(matrix[i][j], i, j+1))) + 1
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
