package main

import (
	"fmt"
)

func main() {
	// common.MakeTestCases("hasAlternatingBits")
	input1 := 5
	fmt.Printf("Expect1: True\n")
	fmt.Printf("Output1: %v\n", hasAlternatingBits(input1))
	input2 := 7
	fmt.Printf("Expect2: False\n")
	fmt.Printf("Output2: %v\n", hasAlternatingBits(input2))
	input3 := 11
	fmt.Printf("Expect3: False\n")
	fmt.Printf("Output3: %v\n", hasAlternatingBits(input3))
	input4 := 42
	fmt.Printf("Expect4: True\n")
	fmt.Printf("Output4: %v\n", hasAlternatingBits(input4))
}

func hasAlternatingBits(n int) bool {
	if n < 2 {
		return true
	}
	s := fmt.Sprintf("%b", n)
	if s[0:1] == s[1:2] {
		return false
	}
	if len(s)%2 != 0 {
		s = "0" + s
	}
	for i := 2; i < len(s); i = i + 2 {
		if s[i:i+2] != s[0:2] {
			return false
		}
	}
	return true
}
