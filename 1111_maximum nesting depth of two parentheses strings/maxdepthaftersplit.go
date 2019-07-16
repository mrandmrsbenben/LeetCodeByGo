package main

import "fmt"

func main() {
	// seq := "(()())"
	seq := "()(())()"
	fmt.Println("Output:", maxDepthAfterSplit(seq))
}

//执行用时 :60 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :7.9 MB, 在所有 Go 提交中击败了100.00%的用户
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
