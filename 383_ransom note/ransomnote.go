package main

import (
	"fmt"
	"strings"
)

func main() {
	ransomNote := ""
	magazine := "a"
	fmt.Printf("Output: %v\n", canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	for i := 0; i < len(ransomNote); i++ {
		if strings.Count(ransomNote, ransomNote[i:i+1]) >
			strings.Count(magazine, ransomNote[i:i+1]) {
			return false
		}
	}
	return true
}
