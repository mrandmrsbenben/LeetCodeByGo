package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abba"
	t := "abab"
	fmt.Printf("Output: %v\n", isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		if strings.Index(s[i+1:], s[i:i+1]) != strings.Index(t[i+1:], t[i:i+1]) {
			return false
		}
	}
	return true
}
func isIsomorphicSlow(s string, t string) bool {
	ms := make(map[string]int)
	mt := make(map[string]int)
	var is, it int
	var ok bool
	for i := 0; i < len(s); i++ {
		if is, ok = ms[s[i:i+1]]; !ok {
			is = -1
		}
		if it, ok = mt[t[i:i+1]]; !ok {
			it = -1
		}
		if is != it {
			return false
		}
		ms[s[i:i+1]] = i
		mt[t[i:i+1]] = i
	}
	return true
}
