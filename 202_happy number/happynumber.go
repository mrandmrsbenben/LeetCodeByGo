package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Output: %v\n", isHappy(9))
}

func isHappy(n int) bool {
	str := strconv.Itoa(n)
	sum := 0
	value := 0
	for i := range str {
		value, _ = strconv.Atoi(str[i : i+1])
		sum = sum + value*value
	}
	fmt.Println(sum)
	if sum >= 10 {
		return isHappy(sum)
	}
	return sum == 1 || sum == 7
}
