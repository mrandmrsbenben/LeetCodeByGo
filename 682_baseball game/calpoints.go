package main

import (
	"fmt"
	"strconv"
)

func main() {
	// common.MakeTestCases("calPoints")
	// str1 := []string{"5", "2", "C", "D", "+"}
	// fmt.Printf("Output: %v\n", calPoints(str1))
	// fmt.Printf("Expect: 30\n")
	// str2 := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	// fmt.Printf("Output: %v\n", calPoints(str2))
	// fmt.Printf("Expect: 27\n")
	str3 := []string{"36", "28", "70", "65", "C", "+", "33", "-46", "84", "C"}
	fmt.Printf("Output: %v\n", calPoints(str3))
	fmt.Printf("Expect: 219\n")
}

func calPoints(ops []string) int {
	p, sum := 0, 0
	last := make([]int, 0)
	for _, r := range ops {
		switch r {
		case "C":
			if len(last) > 0 {
				p = -1 * last[len(last)-1]
				last = last[0 : len(last)-1]
			}
		case "D":
			if len(last) > 0 {
				p = last[len(last)-1] * 2
				last = append(last, p)
			}
		case "+":
			if len(last) > 1 {
				p = last[len(last)-1] + last[len(last)-2]
				last = append(last, p)
			}
		default:
			p, _ = strconv.Atoi(r)
			last = append(last, p)
		}
		sum = sum + p
	}
	return sum
}
