package main

import (
	"fmt"
	"strings"
)

func main() {
	// text := "alice is a good girl she is a good student"
	// first := "a"
	// second := "good"
	text := "we will we will rock you"
	first := "we"
	second := "will"
	fmt.Printf("Output: %v\n", findOcurrences(text, first, second))
}

//执行用时 :0 ms, 在所有Go提交中击败了100.00%的用户
//内存消耗 :2.1 MB, 在所有Go提交中击败了100.00%的用户
func findOcurrences(text string, first string, second string) []string {
	ocurs := make([]string, 0)
	s := strings.Fields(text)
	for i := 0; i < len(s)-2; i++ {
		if s[i] == first && s[i+1] == second {
			ocurs = append(ocurs, s[i+2])
		}
	}
	return ocurs
}
