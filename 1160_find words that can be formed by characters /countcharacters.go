package main

import (
	"fmt"
	"strings"
)

func main() {
	// words := []string{"cat", "bt", "hat", "tree"}
	// chars := "atach"
	words := []string{"hello", "world", "leetcode"}
	chars := "welldonehoneyr"
	fmt.Println("Output:", countCharacters(words, chars))
}

//执行用时 :28 ms, 在所有 Go 提交中击败了75.00%的用户
//内存消耗 :6.3 MB, 在所有 Go 提交中击败了100.00%的用户
func countCharacters(words []string, chars string) int {
	cnt := 0

	// make character's map
	cm := make(map[rune]int)
	for _, c := range chars {
		if n, ok := cm[c]; ok {
			cm[c] = n + 1
		} else {
			cm[c] = 1
		}
	}

	// find which words can be formed
	var cbf bool
	for _, w := range words {
		cbf = true
		for _, c := range w {
			if n, ok := cm[c]; ok {
				if strings.Count(w, string(c)) > n {
					cbf = false
					break
				}
			} else {
				cbf = false
				break
			}
		}
		if cbf {
			cnt += len(w)
		}
	}

	return cnt
}
