package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "FlaG"
	fmt.Printf("Output: %v\n", detectCapitalUse(word))
}

func detectCapitalUse(word string) bool {
	return word == strings.ToUpper(word) || word == strings.ToLower(word) ||
		word == strings.Title(strings.ToLower(word))
}
