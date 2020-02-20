package main

import "fmt"

func main() {
	fmt.Println("Output: ", numWays(100))
}

func numWays(n int) int {
	if n <= 1 {
		return 1
	}

	cur := 0
	prev1, prev2 := 1, 1

	for i := 2; i <= n; i++ {
		cur = (prev1 + prev2) % 1000000007
		prev2, prev1 = prev1, cur
	}
	return cur
}
