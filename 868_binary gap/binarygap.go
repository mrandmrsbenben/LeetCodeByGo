package main

import (
	"fmt"
	"strings"
)

func main() {
	// common.MakeTestCases("binaryGap")
	input1 := 22
	fmt.Printf("Input :%b\n", input1)
	fmt.Printf("Expect1: 2\n")
	fmt.Printf("Output1: %v\n", binaryGap(input1))
	input2 := 5
	fmt.Printf("Input :%b\n", input2)
	fmt.Printf("Expect2: 2\n")
	fmt.Printf("Output2: %v\n", binaryGap(input2))
	input3 := 6
	fmt.Printf("Input :%b\n", input3)
	fmt.Printf("Expect3: 1\n")
	fmt.Printf("Output3: %v\n", binaryGap(input3))
	input4 := 8
	fmt.Printf("Input :%b\n", input4)
	fmt.Printf("Expect4: 0\n")
	fmt.Printf("Output4: %v\n", binaryGap(input4))
	input5 := 20
	fmt.Printf("Input :%b\n", input5)
	fmt.Printf("Expect5: 2\n")
	fmt.Printf("Output5: %v\n", binaryGap(input5))
}

func binaryGap(N int) int {
	s := fmt.Sprintf("%b", N)
	if strings.Count(s, "1") < 2 {
		return 0
	}
	s = s[0 : strings.LastIndex(s, "1")+1]
	str := strings.Split(s, "1")
	gap := 0
	for _, v := range str {
		if len(v) > gap {
			gap = len(v)
		}
	}
	return gap + 1
}
