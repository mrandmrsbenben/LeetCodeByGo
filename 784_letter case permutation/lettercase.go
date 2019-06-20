package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// common.MakeTestCases("letterCasePermutation")
	S1 := "a1b2"
	fmt.Printf("Expect1: {\"a1b2\", \"a1B2\", \"A1b2\", \"A1B2\"}\n")
	fmt.Printf("Output1: %v\n", letterCasePermutation(S1))
	S2 := "3z4"
	fmt.Printf("Expect2: {\"3z4\", \"3Z4\"}\n")
	fmt.Printf("Output2: %v\n", letterCasePermutation(S2))
	S3 := "12345"
	fmt.Printf("Expect3: {\"12345\"}\n")
	fmt.Printf("Output3: %v\n", letterCasePermutation(S3))
}

func letterCasePermutation(S string) []string {
	if len(S) == 0 {
		return []string{S}
	}
	strs := make([]string, 0)
	if unicode.IsLetter(rune(S[0])) {
		strs = append(strs, strings.ToLower(S[0:1]))
		strs = append(strs, strings.ToUpper(S[0:1]))
	} else {
		strs = append(strs, S[0:1])
	}
	for i := 1; i < len(S); i++ {
		if unicode.IsLetter(rune(S[i])) {
			str2 := make([]string, len(strs))
			copy(str2, strs)
			for j := range strs {
				strs[j] = strs[j] + strings.ToLower(S[i:i+1])
				str2[j] = str2[j] + strings.ToUpper(S[i:i+1])
			}
			strs = append(strs, str2...)
		} else {
			for j := range strs {
				strs[j] = strs[j] + S[i:i+1]
			}
		}
	}
	return strs
}
