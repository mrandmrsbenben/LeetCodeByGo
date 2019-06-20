package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "leetcode"
	fmt.Printf("Output: %v\n", reverseVowels(s))
}

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	v := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if strings.Count(vowels, s[i:i+1]) != 0 {
			v = append(v, s[i:i+1])
		}
	}
	str := ""
	for i := 0; i < len(s); i++ {
		if strings.Count(vowels, s[i:i+1]) != 0 {
			str = str + v[len(v)-1]
			v = v[0 : len(v)-1]
		} else {
			str = str + s[i:i+1]
		}
	}
	return str
}
