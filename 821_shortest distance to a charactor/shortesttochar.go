package main

import (
	"fmt"
	"strings"
)

func main() {
	S := "loveleetcode"
	C := 'l'
	fmt.Printf("Output:%d\n", shortestToChar(S, byte(C)))
}

func shortestToChar(S string, C byte) []int {
	ds := make([]int, len(S))
	var l, r int
	for i := range S {
		if S[i:i+1] == string(C) {
			ds[i] = 0
		} else {
			l = strings.LastIndex(S[0:i], string(C))
			if l != -1 {
				l = i - l
			}
			r = strings.IndexByte(S[i:], C)
			if l < r {
				if l != -1 {
					ds[i] = l
				} else {
					ds[i] = r
				}
			} else if l > r {
				if r != -1 {
					ds[i] = r
				} else {
					ds[i] = l
				}
			} else {
				ds[i] = r
			}
		}
	}
	return ds
}
