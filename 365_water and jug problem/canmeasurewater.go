package main

import "fmt"

func main() {
	x := 2
	y := 6
	z := 5
	fmt.Println("Output: ", canMeasureWater(x, y, z))
}

func canMeasureWater(x int, y int, z int) bool {
	if x == 0 || y == 0 {
		return z == 0
	}
	if x+y == z {
		return true
	} else if x+y < z {
		return false
	}
	gcd := func(a, b int) int {
		for a%b != 0 {
			a, b = b, a%b
		}
		return b
	}
	return z%gcd(x, y) == 0
}
