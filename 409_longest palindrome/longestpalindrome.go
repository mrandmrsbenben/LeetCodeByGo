package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abbbcccddd"
	fmt.Printf("Output: %d\n", longestPalindrome(s))
}

func longestPalindrome(s string) int {
	count := 0
	m := make(map[string]int)
	hasOdd := false
	c := 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i:i+1]]; ok {
			continue
		}
		c = strings.Count(s, s[i:i+1])
		if c%2 == 0 {
			count = count + c
		} else {
			count = count + c/2*2
			hasOdd = true
		}
		m[s[i:i+1]] = 1
	}
	if hasOdd {
		count = count + 1
	}
	return count
}
