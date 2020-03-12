package main

import (
	"fmt"
)

func main() {
	// s := "aa"
	// p := "a"
	// s := "aa"
	// p := "a*"
	// s := "ab"
	// p := ".bb*"
	// s := "abbcde"
	// p := "..b.*e"
	// s := "mississippi"
	// p := "mis*is*p*."
	s := "aaca"
	p := "ab*a*c*a"
	fmt.Println("Output: ", isMatch(s, p))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.9 MB, 在所有 Go 提交中击败了19.88%的用户
func isMatch(s string, p string) bool {
	res := make([][]int, len(s)+1)
	for i := range res {
		res[i] = make([]int, len(p)+1)
	}

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if res[i][j] != 0 {
			return res[i][j] == 1
		}
		ans := false
		if j == len(p) {
			ans = i == len(s)
		} else {
			first := i < len(s) && (p[j] == s[i] || p[j] == '.')
			if j+1 < len(p) && p[j+1] == '*' {
				ans = dp(i, j+2) || (first && dp(i+1, j))
			} else {
				ans = first && dp(i+1, j+1)
			}
		}
		if ans {
			res[i][j] = 1
		} else {
			res[i][j] = -1
		}
		return ans
	}
	return dp(0, 0)
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.4 MB, 在所有 Go 提交中击败了50.00%的用户
func isMatchV2(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			first := i < len(s) && (p[j] == s[i] || p[j] == '.')
			if j+1 < len(p) && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (first && dp[i+1][j])
			} else {
				dp[i][j] = first && dp[i+1][j+1]
			}
		}
	}

	return dp[0][0]
}

//执行用时 :20 ms, 在所有 Go 提交中击败了29.14%的用户
//内存消耗 :2.2 MB, 在所有 Go 提交中击败了82.32%的用户
func isMatchV1(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	first := !(s == "") && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (first && isMatch(s[1:], p))
	}
	return first && isMatch(s[1:], p[1:])
}
