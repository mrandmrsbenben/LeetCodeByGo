package main

import "fmt"

func main() {
	fmt.Println("Output: ", integerBreak(3))
}

func integerBreak(n int) int {
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
