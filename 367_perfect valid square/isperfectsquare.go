package main

import "fmt"

func main() {
	fmt.Printf("Output: %v\n", isPerfectSquare(256*256))
}

// 执行用时 : 0 ms, 在Valid Perfect Square的Go提交中击败了100.00% 的用户
// 内存消耗 : 1.9 MB, 在Valid Perfect Square的Go提交中击败了93.33% 的用户
func isPerfectSquare(num int) bool {
	if num%2 == 0 {
		for i := 0; i*i <= num; i = i + 2 {
			if i*i == num {
				return true
			}
		}
	} else {
		for i := 1; i*i <= num; i = i + 2 {
			if i*i == num {
				return true
			}
		}
	}
	return false
}
