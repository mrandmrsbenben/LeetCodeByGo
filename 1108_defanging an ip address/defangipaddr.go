package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Output:", defangIPaddr("1.1.1.1"))
	fmt.Println("Output:", defangIPaddr("255.100.50.0"))
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Defanging an IP Address.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Defanging an IP Address.
func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
