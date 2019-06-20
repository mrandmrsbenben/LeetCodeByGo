package main

import (
	"fmt"
	"strings"
)

func main() {
	A1 := []string{"bella", "label", "roller"}
	A2 := []string{"fill", "lock", "cook"}
	fmt.Printf("Output:%s\n", commonChars(A1))
	fmt.Printf("Output:%s\n", commonChars(A2))
}

func commonChars(A []string) []string {
	com := make(map[string]int)
	word := A[0]
	chars := []byte(word)
	for _, c := range chars {
		com[string(c)] = strings.Count(word, string(c))
	}
	for i := 1; i < len(A); i++ {
		for ch, cnt := range com {
			n := strings.Count(A[i], ch)
			if n == 0 {
				delete(com, ch)
			} else if n > 0 && n < cnt {
				com[ch] = n
			}
		}
	}
	comc := make([]string, 0)
	for k, v := range com {
		for i := 0; i < v; i++ {
			comc = append(comc, k)
		}
	}
	return comc
}
