package main

import "fmt"

func main() {
	// fmt.Println("Output: ", minDistance("horse", "ros"))
	fmt.Println("Output: ", minDistance("intention", "execution"))
}

func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1)+1, len(word2)+1
	dp := make([][]int, l1)
	for i := 0; i < l1; i++ {
		dp[i] = make([]int, l2)
		dp[i][0] = i
	}
	for j := 0; j < l2; j++ {
		dp[0][j] = j
	}
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[l1-1][l2-1]
}

func min(a, b, c int) int {
	if a < b {
		if c < a {
			return c
		}
		return a
	}
	if c < b {
		return c
	}
	return b
}
