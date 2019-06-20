package main

import (
	"fmt"
	"unicode"
)

func main() {
	// common.MakeTestCases("reverseOnlyLetters")
	input1 := "ab-cd-3"
	fmt.Printf("Expect1: dc-ba-3\n")
	fmt.Printf("Output1: %v\n", reverseOnlyLetters(input1))
	input2 := "a-bC-dEf-ghIj"
	fmt.Printf("Expect2: j-Ih-gfE-dCba\n")
	fmt.Printf("Output2: %v\n", reverseOnlyLetters(input2))
	input3 := "Test1ng-Leet3:=code-Q!"
	fmt.Printf("Expect3: Qedo1ct-eeLg3:=ntse-T!\n")
	fmt.Printf("Output3: %v\n", reverseOnlyLetters(input3))
}

func reverseOnlyLetters(S string) string {
	rs := make([]rune, 0)
	for i := 0; i < len(S); i++ {
		if unicode.IsLetter(rune(S[i])) {
			rs = append(rs, rune(S[i]))
		}
	}
	outS := ""
	for i := 0; i < len(S); i++ {
		if unicode.IsLetter(rune(S[i])) {
			outS = outS + string(rs[len(rs)-1])
			rs = rs[0 : len(rs)-1]
		} else {
			outS = outS + S[i:i+1]
		}
	}
	return outS
}
