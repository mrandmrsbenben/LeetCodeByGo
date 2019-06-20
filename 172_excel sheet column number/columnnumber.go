package main

import (
	"fmt"
	"math"
)

func main() {
	input := []string{"A", "AB", "ZY"}
	for _, s := range input {
		fmt.Printf("Input:%s, Output:%d\n", s, titleToNumber(s))
	}
}

func titleToNumber(s string) int {
	num := 0
	chars := []byte(s)
	for i, c := range chars {
		num = num + int(rune(c)-rune('A')+1)*int(math.Pow(26, float64(len(s)-i-1)))
	}
	return num
}
