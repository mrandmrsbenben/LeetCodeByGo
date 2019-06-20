package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Output: %d\n", climbStairs(1))
}

func climbStairs(n int) int {
	s := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i < 3 {
			s[i] = i
		} else {
			s[i] = s[i-1] + s[i-2]
		}
	}
	return s[n]
}
