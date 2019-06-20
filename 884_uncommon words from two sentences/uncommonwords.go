package main

import (
	"fmt"
	"strings"
)

func main() {
	// A := "this apple is sweet"
	// B := "this apple is sour"
	A := "apple apple"
	B := "banana"
	fmt.Printf("Output: %v\n", uncommonFromSentences(A, B))
	// fmt.Println(len(strings.Fields(A)))
}

func uncommonFromSentences(A string, B string) []string {
	strs := make([]string, 0)
	sa := strings.Fields(A)
	ma := make(map[string]int)
	for _, s := range sa {
		if _, ok := ma[s]; ok {
			ma[s] = 1
		} else {
			ma[s] = 0
		}
	}
	sb := strings.Fields(B)
	for _, s := range sb {
		if _, ok := ma[s]; ok {
			ma[s] = 1
		} else {
			ma[s] = 0
		}
	}
	for s, v := range ma {
		if v == 0 {
			strs = append(strs, s)
		}
	}
	return strs
}
