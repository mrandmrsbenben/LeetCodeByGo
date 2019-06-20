package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 38
	fmt.Printf("Output: %d\n", addDigits(num))
}

func addDigits2(num int) int {
	var n int
	str := strconv.Itoa(num)
	for i := 0; i < len(str); i++ {
		x, _ := strconv.Atoi(str[i : i+1])
		n = n + x
	}
	if n >= 10 {
		return addDigits(n)
	}
	return n
}

func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}
