package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%f\n", math.Pow(2, 31))
	fmt.Printf("Output: %v\n", isUgly(-2147483648))
}

func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	if num%2 == 0 {
		return isUgly(num / 2)
	} else if num%3 == 0 {
		return isUgly(num / 3)
	} else if num%5 == 0 {
		return isUgly(num / 5)
	}
	return num == 1
}
