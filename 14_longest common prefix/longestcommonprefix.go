package main

import (
	"fmt"
	"strings"
)

func main() {
	// common.MakeTestCases("longestCommonPrefix")
	input1 := []string{"flower", "flow", "flight"}
	fmt.Printf("Expect1: fl\n")
	fmt.Printf("Output1: %v\n", longestCommonPrefix(input1))
	input2 := []string{"dog", "racecar", "car"}
	fmt.Printf("Expect2: \n")
	fmt.Printf("Output2: %v\n", longestCommonPrefix(input2))
	input3 := []string{"c", "c"}
	fmt.Printf("Expect3: c\n")
	fmt.Printf("Output3: %v\n", longestCommonPrefix(input3))
}

//执行用时 :0 ms, 在所有Go提交中击败了100.00%的用户
//内存消耗 :2.6 MB, 在所有Go提交中击败了17.20%的用户
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	str := "," + strings.Join(strs, ",")
	s := "," + strs[0]
	var common string
	for i := 2; i <= len(s); i++ {
		if strings.Count(str, s[0:i]) == len(strs) {
			common = s[1:i]
		} else {
			break
		}
	}
	return common
}
