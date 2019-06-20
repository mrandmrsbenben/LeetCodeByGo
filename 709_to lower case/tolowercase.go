package main

import (
	"fmt"
	"strings"
)

func main() {
	testcase := []string{"Hello", "here", "LOVELY", "Amy110", "123"}
	for _, t := range testcase {
		fmt.Printf("%s to Lower Case:%s\n", t, toLowerCase(t))
	}
}

func toLowerCase(str string) string {
	return strings.ToLower(str)
}
