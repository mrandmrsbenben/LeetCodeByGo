package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "tree"
	// s := "cccaaa"
	// s := "Aabb"
	fmt.Println("Output: ", frequencySort(s))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了94.05%的用户
//内存消耗 :4.9 MB, 在所有 Go 提交中击败了100.00%的用户
func frequencySort(s string) string {
	freq := make(map[byte]int)
	for i := range s {
		freq[s[i]]++
	}

	strs := make([]string, len(freq))
	i := 0
	for k, v := range freq {
		strs[i] = strings.Repeat(string(k), v)
		i++
	}
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})

	res := strings.Join(strs, "")
	return res
}

//执行用时 :52 ms, 在所有 Go 提交中击败了36.93%的用户
//内存消耗 :13.9 MB, 在所有 Go 提交中击败了100.00%的用户
func frequencySortV1(s string) string {
	if len(s) == 0 {
		return s
	}
	cnts := make([]map[byte]int, 0)
	cm := make(map[byte]int)
	for i := range s {
		if _, ok := cm[s[i]]; ok {
			cm[s[i]]++
		} else {
			cm[s[i]] = 1
		}
		if len(cnts) < cm[s[i]] {
			if len(cnts) == 0 {
				m := make(map[byte]int)
				m[s[i]] = 1
				cnts = []map[byte]int{m}
			} else {
				delete(cnts[len(cnts)-1], s[i])
				m := make(map[byte]int)
				m[s[i]] = cm[s[i]]
				cnts = append(cnts, m)
			}
		} else {
			if cm[s[i]] > 1 {
				delete(cnts[cm[s[i]]-2], s[i])
			}
			cnts[cm[s[i]]-1][s[i]] = cm[s[i]]
		}
	}

	res := ""
	for i := len(cnts) - 1; i >= 0; i-- {
		for c := range cnts[i] {
			res += strings.Repeat(string(c), i+1)
		}
	}
	return res
}
