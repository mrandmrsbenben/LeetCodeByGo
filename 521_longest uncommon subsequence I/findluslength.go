package main

import "fmt"

func main() {
	a, b := "aba", "cdc"
	fmt.Printf("Output: %d\n", findLUSlength(a, b))
}

func findLUSlength(a string, b string) int {
	if len(a) > len(b) {
		return len(a)
	} else if a == b {
		return -1
	}
	return len(b)
}
