package main

import "fmt"

func main() {
	// common.MakeTestCases("largeGroupPositions")
	input1 := "abbxxxxzzy"
	fmt.Printf("Expect1: {{3,6}}\n")
	fmt.Printf("Output1: %v\n", largeGroupPositions(input1))
	input2 := "abc"
	fmt.Printf("Expect2: {}\n")
	fmt.Printf("Output2: %v\n", largeGroupPositions(input2))
	input3 := "abcdddeeeeaabbbcd"
	fmt.Printf("Expect3: {{3,5},{6,9},{12,14}}\n")
	fmt.Printf("Output3: %v\n", largeGroupPositions(input3))
	input4 := "bbbbb"
	fmt.Printf("Output4: %v\n", largeGroupPositions(input4))
}

/*
执行用时 : 4 ms, 在Positions of Large Groups的Go提交中击败了95.45% 的用户
内存消耗 : 6.1 MB, 在Positions of Large Groups的Go提交中击败了50.00% 的用户
*/
func largeGroupPositions(S string) [][]int {
	p := [][]int{}
	if len(S) < 3 {
		return p
	}
	start := 0
	for i := 1; i < len(S); i++ {
		if S[i] != S[start] {
			if i-1-start >= 2 {
				p = append(p, []int{start, i - 1})
			}
			start = i
		} else if i == len(S)-1 && i-start >= 2 {
			p = append(p, []int{start, i})
		}
	}
	return p
}
