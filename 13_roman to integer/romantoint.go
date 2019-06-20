package main

import (
	"fmt"
)

func main() {
	// common.MakeTestCases("romanToInt")
	input1 := "III"
	fmt.Printf("Expect1: 3\n")
	fmt.Printf("Output1: %v\n", romanToInt(input1))
	input2 := "MCMXCIV"
	fmt.Printf("Expect2: 1994\n")
	fmt.Printf("Output2: %v\n", romanToInt(input2))
	input3 := "IX"
	fmt.Printf("Expect3: 9\n")
	fmt.Printf("Output3: %v\n", romanToInt(input3))
	input4 := "LVIII"
	fmt.Printf("Expect4: 58\n")
	fmt.Printf("Output4: %v\n", romanToInt(input4))
	input5 := "MCDLXXVI"
	fmt.Printf("Expect5: 1476\n")
	fmt.Printf("Output5: %v\n", romanToInt(input5))

}

func romanToInt(s string) int {
	m := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900}
	if len(s) == 1 {
		return m[s]
	}
	if len(s) == 2 {
		val, ok := m[s]
		if !ok {
			val = m[s[0:1]] + m[s[1:]]
		}
		return val
	}
	var str string
	var sum int
	for i := 0; i < len(s); i++ {
		if i+2 > len(s)-1 {
			str = s[i:]
		} else {
			str = s[i : i+2]
		}
		val, ok := m[str]
		if !ok {
			sum = sum + m[str[0:1]]
			if _, ok = m[str[1:]]; ok {
				continue
			}
		} else {
			sum = sum + val
			i = i + 1
		}
	}
	return sum
}
