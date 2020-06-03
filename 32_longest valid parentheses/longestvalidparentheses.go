package main

import "fmt"

func main() {
	// s := "())((())()))"
	s := "()(()"
	fmt.Println("Output: ", longestValidParentheses(s))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :4.3 MB, 在所有 Go 提交中击败了100.00%的用户
func longestValidParentheses(s string) int {
	res := 0
	dp := make([]int, len(s)+1)
	lcnt, rcnt := make([]int, len(s)), make([]int, len(s))
	cnt := 0
	for i := range s {
		lcnt[i] = cnt
		if s[i] == '(' {
			cnt++
		} else if cnt > 0 {
			cnt--
		}
	}

	cnt = 0
	for i := len(s) - 1; i >= 0; i-- {
		rcnt[i] = cnt
		if s[i] == ')' {
			cnt++
		} else if cnt > 0 {
			cnt--
		}
	}

	for i := range s {
		if s[i] == '(' {
			if rcnt[i] > 0 {
				dp[i+1] = dp[i]
			} else {
				dp[i+1] = 0
			}
		} else {
			if lcnt[i] > 0 {
				dp[i+1] = dp[i] + 2
			} else {
				dp[i+1] = 0
			}
		}
		if dp[i+1] > res {
			res = dp[i+1]
		}
	}

	return res
}
