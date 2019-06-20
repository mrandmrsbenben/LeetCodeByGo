package main

import (
	"fmt"
	"math"
)

func main() {
	area := 10000000
	fmt.Printf("Output: %v\n", constructRectangle(area))
}

func constructRectangle(area int) []int {
	sqrt := int(math.Round(math.Sqrt(float64(area))))
	l, w := area, 1
	for i := sqrt; i >= 0; i-- {
		if area%i == 0 {
			l, w = area/i, i
			break
		}
	}
	return []int{l, w}
}
