package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// common.MakeTestCases("bitwiseComplement")
	input1 := 5
	fmt.Printf("Expect1: 2\n")
	fmt.Printf("Output1: %v\n", bitwiseComplement(input1))
	input2 := 7
	fmt.Printf("Expect2: 0\n")
	fmt.Printf("Output2: %v\n", bitwiseComplement(input2))
	input3 := 10
	fmt.Printf("Expect3: 5\n")
	fmt.Printf("Output3: %v\n", bitwiseComplement(input3))

}

func bitwiseComplement(N int) int {
	s := strconv.FormatInt(int64(N), 2)
	s = strings.Replace(s, "0", "*", -1)
	s = strings.Replace(s, "1", "0", -1)
	s = strings.Replace(s, "*", "1", -1)
	i, _ := strconv.ParseInt(s, 2, 64)
	return int(i)
}
