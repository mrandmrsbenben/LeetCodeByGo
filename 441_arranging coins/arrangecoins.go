package main

import "fmt"

func main() {
	fmt.Printf("Output: %d\n", arrangeCoins(8))
}

// 执行用时 : 8 ms, 在Arranging Coins的Go提交中击败了89.74% 的用户
// 内存消耗 : 2.3 MB, 在Arranging Coins的Go提交中击败了40.74% 的用户
func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}
	steps := 1
	for {
		n -= steps
		if n <= steps {
			break
		}
		steps++
	}
	return steps
}
