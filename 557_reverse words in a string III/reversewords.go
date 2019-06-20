package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Let's take LeetCode contest"
	expect := "s'teL ekat edoCteeL tsetnoc"
	output := reverseWords(input)
	fmt.Printf("%v Output:%s\n", output == expect, output)
}

func reverseWords(s string) string {
	strs := strings.Fields(s)
	output := ""
	for _, str := range strs {
		b := []byte(str)
		for i := 0; i < len(b)/2; i++ {
			b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
		}
		if len(output) == 0 {
			output = string(b)
		} else {
			output = output + " " + string(b)
		}
	}
	return output
}
