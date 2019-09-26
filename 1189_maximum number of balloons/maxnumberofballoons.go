package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Output: ", maxNumberOfBalloons("nlaebolko"))
	fmt.Println("Output: ", maxNumberOfBalloons("loonbalxballpoon"))
	fmt.Println("Output: ", maxNumberOfBalloons("leetcode"))
}

func maxNumberOfBalloons(text string) int {
	cnt := []int{strings.Count(text, "b"),
		strings.Count(text, "a"),
		strings.Count(text, "l") / 2,
		strings.Count(text, "o") / 2,
		strings.Count(text, "n")}
	sort.Ints(cnt)
	return cnt[0]
}
