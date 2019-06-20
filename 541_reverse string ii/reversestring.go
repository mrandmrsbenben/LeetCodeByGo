package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcdefg"
	k := 2
	// fmt.Printf("%v, %d\n", strings.SplitAfter(s, ""), len(strings.SplitAfter(s, "")))
	fmt.Printf("Output: %v\n", reverseStr(s, k))
}

func reverseStr(s string, k int) string {
	str := strings.SplitAfter(s, "")
	j := 0
	swaptoend := false
	for i := 0; i < len(str); i++ {
		if i%(2*k) == 0 {
			if len(str)-i < k {
				swaptoend = true
			}
			j = 0
		} else {
			j = j + 1
		}
		if swaptoend && len(str)-1-i > j {
			str[i], str[len(str)-1-j] = str[len(str)-1-j], str[i]
		} else if !swaptoend && j < k/2 {
			str[i], str[i+k-1-i%k-j] = str[i+k-1-i%k-j], str[i]
		}
		// if i+k > len(str) && i%k < k/2 && i%(2*k) < k {
		// 	str[i], str[len(str)-1-i%(2*k)] = str[len(str)-1-i%(2*k)], str[i]
		// } else {
		// 	if i%(2*k) < k && i%k < k/2 {
		// 		str[i], str[i+k-i%k*2-1] = str[i+k-i%k*2-1], str[i]
		// 	}
		// }
	}
	return strings.Join(str, "")
}

func reverseStrSlow(s string, k int) string {
	rs := ""
	substr := ""
	for i := 0; i < len(s); i = i + 2*k {
		if i+2*k < len(s) {
			substr = s[i : i+2*k]
		} else {
			substr = s[i:]
		}
		if len(substr) < k {
			for j := len(substr) - 1; j >= 0; j-- {
				rs = rs + substr[j:j+1]
			}
		} else {
			for j := 0; j < k; j++ {
				rs = rs + substr[k-1-j:k-j]
			}
			rs = rs + substr[k:]
		}
	}
	return rs
}
