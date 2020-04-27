package main

import "fmt"

func main() {
	text1 := "abcde"
	text2 := "acce"
	// text1 := "abc"
	// text2 := "def"
	// text1 := "ezupkr"
	// text2 := "ubmrapg"
	fmt.Println("Output: ", longestCommonSubsequence(text1, text2))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :10.4 MB, 在所有 Go 提交中击败了100.00%的用户
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
