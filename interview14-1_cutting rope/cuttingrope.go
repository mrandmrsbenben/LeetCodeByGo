package main

import "fmt"

func main() {
	fmt.Println("Output: ", cuttingRope(3))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
func cuttingRope(n int) int {
	c := 2
	val, max := 0, 0
	for {
		val = breakValue(n, c)
		if val >= max {
			max = val
			c++
		} else {
			break
		}
	}
	return max
}

func breakValue(n, c int) int {
	val := 1
	for c > 1 {
		x := n / c
		val *= x
		n -= x
		c--
	}
	val *= n
	return val
}
