package main

import (
	"fmt"
	"strings"
)

func main() {
	// paragraph := "abc abc? abcd the jeff!"
	// banned := []string{"abc", "abcd", "jeff"}
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	fmt.Println("Output:", mostCommonWord(paragraph, banned))
}

//执行用时 :0 ms, 在所有Go提交中击败了100.00%的用户
//内存消耗 :3.1 MB, 在所有Go提交中击败了66.67%的用户
func mostCommonWord(paragraph string, banned []string) string {
	paragraph = " " + strings.ToLower(paragraph) + " "
	symbol := []string{"!", "?", "'", ",", ";", "."}
	for _, s := range symbol {
		paragraph = strings.Replace(paragraph, s, " ", -1)
	}
	bm := make(map[string]int)
	for _, s := range banned {
		bm[s] = 1
	}
	words := strings.Fields(paragraph)
	max := 0
	var mcw string
	cm := make(map[string]int)
	for _, w := range words {
		if _, ok := bm[w]; ok {
			continue
		}
		if _, ok := cm[w]; ok {
			cm[w] = cm[w] + 1
		} else {
			cm[w] = 1
		}
		if cm[w] > max {
			max = cm[w]
			mcw = w
		}
	}
	return mcw
}
