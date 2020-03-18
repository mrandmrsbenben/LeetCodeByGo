package main

import "fmt"

// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
func main() {
	fmt.Println("Output: ", permutation("abaa"))
}

//执行用时 :68 ms, 在所有 Go 提交中击败了72.95%的用户
//内存消耗 :6.9 MB, 在所有 Go 提交中击败了100.00%的用户
func permutation(s string) []string {
	if len(s) == 0 {
		return []string{}
	} else if len(s) == 1 {
		return []string{s}
	} else if len(s) == 2 {
		if s[0] == s[1] {
			return []string{s}
		}
		return []string{s, s[1:] + s[0:1]}
	}

	var res, subs []string
	var str string
	m := make(map[rune]int)
	for i, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = 1
			str = s[0:i] + s[i+1:]
			subs = permutation(str)
			for j := range subs {
				subs[j] = string(c) + subs[j]
			}
			res = append(res, subs...)
		}
	}
	return res
}
