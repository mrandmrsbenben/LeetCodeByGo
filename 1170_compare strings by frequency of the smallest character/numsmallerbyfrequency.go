package main

import (
	"fmt"
	"sort"
)

func main() {
	queries := []string{"cbd"}
	words := []string{"zaaaz"}
	// queries := []string{"bbb", "cc"}
	// words := []string{"a", "aa", "aaa", "aaaa"}
	fmt.Println(numSmallerByFrequency(queries, words))
}

//执行用时 :24 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :7.5 MB, 在所有 Go 提交中击败了100.00%的用户
func numSmallerByFrequency(queries []string, words []string) []int {
	wordNums := make([]int, len(words))
	for i := range words {
		wordNums[i] = f(words[i])
	}
	sort.Ints(wordNums)
	l := len(words)
	nums := make([]int, len(queries))
	var index, freq int
	for i := range queries {
		freq = f(queries[i])
		index = sort.SearchInts(wordNums, freq+1)
		nums[i] = l - index
	}
	return nums
}

func f(s string) int {
	minc := 'z' + 1
	cnt := 0
	for _, c := range s {
		if minc > c {
			minc = c
			cnt = 1
		} else if minc == c {
			cnt++
		}
	}
	return cnt
}
