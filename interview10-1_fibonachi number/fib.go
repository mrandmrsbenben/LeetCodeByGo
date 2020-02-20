package main

import "fmt"

func main() {
	fmt.Println("Output: ", fib(5))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func fib(n int) int {
	if n <= 1 {
		return n
	} else if n == 2 {
		return 1
	}

	current := 0
	prev1 := 1
	prev2 := 1

	for i := 3; i <= n; i++ {
		current = (prev1 + prev2) % 1000000007
		prev2, prev1 = prev1, current
	}
	return current
}
