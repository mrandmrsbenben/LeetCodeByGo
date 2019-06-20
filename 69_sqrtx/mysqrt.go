package main

import "fmt"

func main() {
	fmt.Printf("Output: %d\n", mySqrt(2143723787))
}

//执行用时 :8 ms, 在所有Go提交中击败了77.33%的用户
//内存消耗 :2.2 MB, 在所有Go提交中击败了72.55%的用户
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	n := x / 2
	for {
		if n*n <= x {
			break
		}
		n = n / 2
	}
	for {
		if n*n > x {
			return n - 1
		}
		n++
	}
}
