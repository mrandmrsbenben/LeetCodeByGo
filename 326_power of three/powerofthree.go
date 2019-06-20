package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Output: %v\n", isPowerOfThree(27))
}

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	s := strconv.FormatInt(int64(n), 3)
	return s[0:1] == "1" && strings.Count(s, "0") == len(s)-1
}
