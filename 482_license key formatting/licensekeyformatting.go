package main

import (
	"fmt"
	"strings"
)

func main() {
	// S := "5F3Z-2e-9-w"
	// K := 4
	S := "2-5g-3-J"
	K := 2
	fmt.Printf("Output: %v\n", licenseKeyFormatting(S, K))
}

// 执行用时 :252 ms, 在所有Go提交中击败了57.58%的用户
// 内存消耗 :8.1 MB, 在所有Go提交中击败了40.00%的用户
func licenseKeyFormatting(S string, K int) string {
	S = strings.ToUpper(strings.Replace(S, "-", "", -1))
	if len(S) <= K {
		return S
	}
	s := ""
	for i := len(S) - 1; i >= 0; i = i - K {
		if i-K < 0 {
			s = S[0:i+1] + "-" + s
		} else {
			s = S[i-K+1:i+1] + "-" + s
		}
	}
	return s[0 : len(s)-1]
}
