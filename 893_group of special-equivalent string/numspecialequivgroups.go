package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// common.MakeTestCases("numSpecialEquivGroups")
	input1 := []string{"a", "b", "c", "a", "c", "c"}
	// fmt.Println(rune(input1[0][0]) ^ rune(input1[3][0]))
	fmt.Printf("Output: %v\n", numSpecialEquivGroups(input1))
	fmt.Printf("Expect: 3\n")
	input2 := []string{"aa", "bb", "ab", "ba"}
	fmt.Printf("Output: %v\n", numSpecialEquivGroups(input2))
	fmt.Printf("Expect: 4\n")
	input3 := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	fmt.Printf("Output: %v\n", numSpecialEquivGroups(input3))
	fmt.Printf("Expect: 3\n")
	input4 := []string{"abcd", "cdab", "adcb", "cbad"}
	fmt.Printf("Output: %v\n", numSpecialEquivGroups(input4))
	fmt.Printf("Expect: 1\n")
}

func numSpecialEquivGroups(A []string) int {
	groups := make(map[string]int)
	var eKey, oKey []string
	var key string
	for _, s := range A {
		eKey = []string{}
		oKey = []string{}
		for i := 0; i < len(s); i++ {
			if i%2 == 0 {
				eKey = append(eKey, s[i:i+1])
			} else {
				oKey = append(oKey, s[i:i+1])
			}
		}
		sort.Strings(eKey)
		sort.Strings(oKey)
		key = strings.Join(eKey, "") + "&" + strings.Join(oKey, "")
		if _, ok := groups[key]; !ok {
			groups[key] = 0
		}
	}
	return len(groups)
}
