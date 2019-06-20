package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"a"}
	order := "abcdefghijklmnopqrstuvwxyz"
	fmt.Printf("Output: %v\n", isAlienSorted(words, order))
}

func isAlienSorted(words []string, order string) bool {
	compare := func(s1, s2 string) bool {
		for i := 0; i < len(s1); i++ {
			if i >= len(s2) {
				return false
			}
			if s1[i:i+1] == s2[i:i+1] {
				continue
			}
			return strings.Index(order, s1[i:i+1]) < strings.Index(order, s2[i:i+1])
		}
		return true
	}
	for i := 0; i < len(words)-1; i++ {
		if !compare(words[i], words[i+1]) {
			return false
		}
	}
	return true
}
