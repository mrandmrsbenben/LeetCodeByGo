package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Output: %d\n", hammingWeight(11))
}

func hammingWeight(num uint32) int {
	return strings.Count(fmt.Sprintf("%b", num), "1")
}
