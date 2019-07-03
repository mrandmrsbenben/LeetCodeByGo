package main

import (
	"fmt"
	"strings"
)

func main() {
	S := "ababcbacadefegdehijhklij"
	fmt.Println("Output:", partitionLabelsV2(S))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.4 MB, 在所有 Go 提交中击败了33.33%的用户
func partitionLabels(S string) []int {
	ps := make([]int, 0)
	s := make(map[byte]int)
	begin := 0
	substr := ""
	for i := range S {
		if _, ok := s[S[i]]; !ok {
			substr += string(S[i])
			s[S[i]] = 0
		}
		if !strings.ContainsAny(S[i+1:], substr) {
			ps = append(ps, i-begin+1)
			begin = i + 1
			substr = ""
			s = make(map[byte]int)
		}
	}
	return ps
}

//执行用时 :4 ms, 在所有 Go 提交中击败了83.33%的用户
//内存消耗 :2.2 MB, 在所有 Go 提交中击败了66.67%的用户
func partitionLabelsV1(S string) []int {
	ps := make([]int, 0)
	begin := 0
	for i := range S {
		if !strings.ContainsAny(S[i+1:], S[begin:i+1]) {
			ps = append(ps, i-begin+1)
			begin = i + 1
		}
	}
	return ps
}
