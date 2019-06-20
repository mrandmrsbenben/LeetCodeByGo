package main

import "fmt"

func main() {
	digits := []int{0}
	fmt.Printf("Output: %v\n", plusOne(digits))
}

// 执行用时 : 0 ms, 在Plus One的Go提交中击败了100.00% 的用户
// 内存消耗 : 2.2 MB, 在Plus One的Go提交中击败了30.57% 的用户
func plusOne(digits []int) []int {
	add := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += add
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
	return digits
}
