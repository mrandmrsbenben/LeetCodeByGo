package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "AAAAAA"
	str2 := "AAA"
	// str1 := "ABCABC"
	// str2 := "ABC"
	// str1 := "ABABAB"
	// str2 := "ABAB"
	// str1 := "NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM"
	// str2 := "NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM"
	// str1 := "TAUXXTAUXXTAUXXTAUXXTAUXX"
	// str2 := "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"
	// str1 := "FBFKOXFBFKOXFBFKOXFBFKOXFBFKOX"
	// str2 := "FBFKOXFBFKOXFBFKOXFBFKOXFBFKOXFBFKOXFBFKOXFBFKOXFBFKOX"
	fmt.Printf("Output: %v\n", gcdOfStrings(str1, str2))
}

// 执行用时 :4 ms, 在所有Go提交中击败了73.68%的用户
// 内存消耗 :2.5 MB, 在所有Go提交中击败了100.00%的用户
func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	if len(str2) == 1 {
		if len(str1) == strings.Count(str1, str2) {
			return str2
		}
	}
	if len(str1)%len(str2) == 0 && isCommonDivisor(str1, str2) {
		return str2
	}
	var str string
	for i := len(str2) - 1; i > 0; i-- {
		if len(str2)%i > 0 || len(str1)%i > 0 {
			continue
		}
		str = str2[0:i]
		if isCommonDivisor(str2, str) && isCommonDivisor(str1, str) {
			return str
		}
	}
	return ""
}

func isCommonDivisor(str, substr string) bool {
	for i := 0; i < len(str); i += len(substr) {
		if i+len(substr) > len(str) || str[i:i+len(substr)] != substr {
			return false
		}
	}
	return true
}
