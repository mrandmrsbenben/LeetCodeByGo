package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := -16
	fmt.Printf("Output: %v\n", isPowerOfTwo(n))
}

func isPowerOfTwo(n int) bool {
	if n < 0 {
		return false
	}
	return strings.Count(strconv.FormatInt(int64(n), 2), "1") == 1
}
