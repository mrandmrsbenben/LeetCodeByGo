package main

import (
	"fmt"
	"strings"
)

func main() {
	// licensePlate := "1s3 456"
	// words := []string{"looks", "pest", "stew", "show"}
	licensePlate := "GrC8950"
	words := []string{"measure", "other", "every", "base", "according", "level", "meeting", "none", "marriage", "rest"}
	fmt.Printf("Output: %s\n", shortestCompletingWord(licensePlate, words))
}

func shortestCompletingWord(licensePlate string, words []string) string {
	str := ""
	lic := strings.ToLower(licensePlate)
	m := make(map[string]int)
	for i := range lic {
		if lic[i:i+1] < "a" || lic[i:i+1] > "z" {
			continue
		}
		m[lic[i:i+1]] = strings.Count(lic, lic[i:i+1])
	}
	var comp bool
	for _, w := range words {
		comp = true
		for k, v := range m {
			if strings.Count(w, k) < v {
				comp = false
				break
			}
		}
		if comp && (str == "" || len(str) > len(w)) {
			str = w
		}
	}
	return str
}
