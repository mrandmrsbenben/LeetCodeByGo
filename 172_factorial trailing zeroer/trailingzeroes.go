package main

import (
	"fmt"
)

func main() {
	// fmt.Printf("Output: %d\n", trailingZeroes(5))
	// fmt.Printf("Output: %d\n", trailingZeroes(25))
	fmt.Printf("Output: %d\n", trailingZeroes(30))
}

func trailingZeroesV2(n int) int {
	digits := []int{1}
	add := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i]*i + add
		if digits[i] > 9 {
			add = digits[i] / 10
			digits[i] = digits[i] % 10
		} else {
			add = 0
		}
	}
	for add > 0 {
		digits = append([]int{add % 10}, digits...)
		add = add / 10
	}
	cnt := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 0 {
			cnt++
		} else {
			break
		}
	}
	return cnt
}

// 执行用时 : 0 ms, 在Factorial Trailing Zeroes的Go提交中击败了100.00% 的用户
// 内存消耗 : 2.1 MB, 在Factorial Trailing Zeroes的Go提交中击败了37.25% 的用户
func trailingZeroes(n int) int {
	cnt := 0
	for n >= 5 {
		cnt += n / 5
		n = n / 5
	}
	return cnt
}
