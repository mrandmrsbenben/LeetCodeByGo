package main

import "fmt"

func main() {
	fmt.Println("Output: ", guessNumber(10))
}

func guess(num int) int {
	switch {
	case num > 6:
		return 1
	case num < 6:
		return -1
	default:
		return 0
	}
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.1 MB, 在所有 Go 提交中击败了100.00%的用户
func guessNumber(n int) int {
	l, r := 1, n
	for l < r {
		n = l + (r-l)/2
		if guess(n) == 0 {
			break
		} else if guess(n) == 1 {
			l = n + 1
		} else {
			r = n - 1
		}
	}
	return n
}
