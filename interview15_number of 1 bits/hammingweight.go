package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Output: ", hammingWeight(9))
}

func hammingWeight(n int) int {
	return strings.Count(strconv.FormatInt(int64(n), 2), "1")
}
