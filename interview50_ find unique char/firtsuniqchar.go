package main

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
func main() {
	fmt.Println("Output: ", string(firstUniqChar("cc")))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了86.39%的用户
//内存消耗 :5.8 MB, 在所有 Go 提交中击败了100.00%的用户
func firstUniqChar(s string) byte {
	res := make([]int, 26)
	for _, c := range s {
		res[c-'a']++
	}
	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		if res[c-'a'] == 1 {
			return byte(c)
		}
	}
	return ' '
}

func firstUniqCharV2(s string) byte {
	for i := range s {
		if strings.Count(s, s[i:i+1]) == 1 {
			return s[i]
		}
	}
	return ' '
}

//执行用时 :20 ms, 在所有 Go 提交中击败了63.91%的用户
//内存消耗 :5.7 MB, 在所有 Go 提交中击败了100.00%的用户
func firstUniqCharV1(s string) byte {
	mc := make(map[byte]int)
	for i := range s {
		if mc[s[i]] == 0 {
			mc[s[i]] = 1
			if strings.LastIndexByte(s, s[i]) == i {
				return s[i]
			}
		}
	}
	return ' '
}
