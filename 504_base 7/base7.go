package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Output: %v\n", convertToBase7(100))
}

func convertToBase7(num int) string {
	return strconv.FormatInt(int64(num), 7)
}
