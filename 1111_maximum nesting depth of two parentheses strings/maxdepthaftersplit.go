package main

import "fmt"

func main() {
	// seq := "(()())"
	seq := "()(())()"
	fmt.Println("Output:", maxDepthAfterSplit(seq))
}

func maxDepthAfterSplit(seq string) []int {
	answer := make([]int, len(seq))
	a, b := 0, 0
	for i := range seq {
		if seq[i] == '(' {
			answer[i] = a
			a = (a + 1) % 2
		} else {
			answer[i] = b
			b = (b + 1) % 2
		}
	}
	return answer
}
