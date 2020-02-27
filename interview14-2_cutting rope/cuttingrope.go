package main

import "fmt"

func main() {
	fmt.Println("Output: ", cuttingRope(1000))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func cuttingRope(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}
	x := n / 3
	val := 1
	switch n % 3 {
	case 0:
		val = breakValue(x)
	case 1:
		val = breakValue(x-1) * 4
	case 2:
		val = breakValue(x) * 2
	}
	return val % 1000000007
}

func breakValue(n int) int {
	val := 1
	for i := 0; i < n; i++ {
		val = (val * 3) % 1000000007
	}
	return val
}
