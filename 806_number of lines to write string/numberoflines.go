package main

import (
	"fmt"
)

func main() {
	// common.MakeTestCases("numberOfLines")
	widths1 := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	S1 := "abcdefghijklmnopqrstuvwxyz"
	fmt.Printf("Output: %v\n", numberOfLines(widths1, S1))
	fmt.Printf("Expect: {3, 60}\n")
	widths2 := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	S2 := "bbbcccdddaaa"
	fmt.Printf("Output: %v\n", numberOfLines(widths2, S2))
	fmt.Printf("Expect: {2, 4}\n")
}

func numberOfLines(widths []int, S string) []int {
	ret := make([]int, 2)
	ret[0] = 1
	for i := 0; i < len(S); i++ {
		ret[1] = ret[1] + widths[int(rune(S[i])-rune('a'))]
		if ret[1] > 100 {
			ret[0] = ret[0] + 1
			ret[1] = widths[int(rune(S[i])-rune('a'))]
		}
	}
	return ret
}
