package main

import (
	"fmt"
	"strconv"
)

// https://leetcode-cn.com/problems/compress-string-lcci/
func main() {
	S := "aabcccccaaa"
	// S := "abbccd"
	fmt.Println("Output: ", compressString(S))
}

func compressString(S string) string {
	if len(S) <= 2 {
		return S
	}
	res := ""

	cnt := 1
	for i := 1; i < len(S); i++ {
		if S[i] != S[i-1] {
			res += S[i-1:i] + strconv.Itoa(cnt)
			if len(res) >= len(S) {
				return S
			}
			cnt = 1
		} else {
			cnt++
		}
	}
	res += string(S[len(S)-1]) + strconv.Itoa(cnt)
	if len(res) < len(S) {
		return res
	}
	return S
}
