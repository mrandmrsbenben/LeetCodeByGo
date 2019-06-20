package main

import (
	"fmt"
)

func main() {
	a := -2
	b := -1
	fmt.Printf("Output: %d\n", getSum(a, b))
}

func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}
	return a
}
