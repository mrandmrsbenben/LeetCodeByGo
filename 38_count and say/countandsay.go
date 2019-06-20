package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Output: %s\n", countAndSay(5))
}

func countAndSay(n int) string {
	s := "1"
	var say, number string
	var start int
	for i := 2; i <= n; i++ {
		number = s[0:1]
		for j := 0; j < len(s); j++ {
			if s[j:j+1] != number {
				say = say + strconv.Itoa(j-start) + number
				start = j
				number = s[j : j+1]
			}
		}
		say = say + strconv.Itoa(len(s)-start) + number
		s = say
		say = ""
		start = 0
	}
	return s
}
