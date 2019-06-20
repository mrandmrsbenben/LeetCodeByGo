package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "PPALLP"
	fmt.Printf("Output: %v\n", checkRecord(s))
}

func checkRecord(s string) bool {
	if strings.Count(s, "A") > 1 || strings.Count(s, "LLL") > 0 {
		return false
	}
	return true
}
