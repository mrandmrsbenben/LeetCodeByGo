package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 65
	fmt.Printf("Output: %v\n", isPowerOfFour(n))
}

func isPowerOfFour(num int) bool {
	if num < 1 {
		return false
	}
	s := strconv.FormatInt(int64(num), 4)
	return strings.Count(s, "1") == 1 && strings.Count(s, "0") == len(s)-1
}
