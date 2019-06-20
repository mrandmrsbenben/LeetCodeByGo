package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Printf("Output:%v\n", findWords2(input))
}

func findWords(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	output := make([]string, 0)
	for _, s := range words {
		if len(s) == 0 {
			continue
		}
		lowerS := strings.ToLower(s)
		row := -1
		oneRow := true
		for j := 0; j < len(lowerS) && oneRow; j++ {
			c := lowerS[j : j+1]
			for r := range rows {
				if strings.Contains(rows[r], c) {
					if row == -1 {
						row = r
					} else if row != r {
						oneRow = false
					}
					break
				}
			}
		}
		if oneRow {
			output = append(output, s)
		}
	}
	return output
}

func findWords2(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	output := make([]string, 0)
	for _, s := range words {
		if len(s) == 0 {
			continue
		}
		lowerS := strings.ToLower(s)
		oneRow := false
		for _, r := range rows {
			if strings.ContainsAny(lowerS, r) {
				oneRow = !oneRow
				if !oneRow {
					break
				}
			}
		}
		if oneRow {
			output = append(output, s)
		}
	}
	return output
}
