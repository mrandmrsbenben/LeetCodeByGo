package main

import (
	"fmt"
	"strings"
)

func main() {
	// common.MakeTestCases("toGoatLatin")
	input1 := "I speak Goat Latin"
	fmt.Printf("Expect1: Imaa peaksmaaa oatGmaaaa atinLmaaaaa\n")
	fmt.Printf("Output1: %v\n", toGoatLatin(input1))
	input2 := "The quick brown fox jumped over the lazy dog"
	fmt.Printf("Expect2: heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa\n")
	fmt.Printf("Output2: %v\n", toGoatLatin(input2))
	// fmt.Println(strings.IndexAny(input1[2:], "aeiouAEIOU"))
}

func toGoatLatin(S string) string {
	fields := strings.Fields(S)
	for i, s := range fields {
		if strings.IndexAny(s[0:1], "aeiouAEIOU") != 0 {
			s = s[1:] + s[0:1]
		}
		fields[i] = s + "ma" + strings.Repeat("a", i+1)
	}
	return strings.Join(fields, " ")
}
