package main

import (
	"fmt"
)

func main() {
	// words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	// pattern := "abb"
	words := []string{"abc", "cba", "xyx", "yxx", "yyx"}
	pattern := "abc"
	fmt.Println("Output:", findAndReplacePattern(words, pattern))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了87.50%的用户
//内存消耗 :3 MB, 在所有 Go 提交中击败了20.00%的用户
func findAndReplacePattern(words []string, pattern string) []string {
	samePattern := make([]string, 0)
	for _, w := range words {
		if transToPattern(w, pattern) == pattern {
			samePattern = append(samePattern, w)
		}
	}
	return samePattern
}

func transToPattern(word, pattern string) string {
	dic := make(map[string]string)
	pm := make(map[string]int)
	var str string
	var c, p string
	for i := 0; i < len(pattern); i++ {
		c = word[i : i+1]
		p = pattern[i : i+1]
		s, ok := dic[c]
		if !ok {
			if _, exist := pm[p]; exist {
				return ""
			}
			pm[p] = 0
			dic[c] = p
			s = p
		}
		str += s
	}
	return str
}
