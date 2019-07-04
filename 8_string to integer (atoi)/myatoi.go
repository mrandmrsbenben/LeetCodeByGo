package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	// fmt.Println("Output:", myAtoi("+42"))
	// fmt.Println("Output:", myAtoi("   -42"))
	// fmt.Println("Output:", myAtoi("4193 with words"))
	// fmt.Println("Output:", myAtoi("words and 987"))
	// fmt.Println("Output:", myAtoi("-91283472332"))
	// fmt.Println("Output:", myAtoi("2147483648"))
	// fmt.Println("Output:", myAtoi(" "))
	fmt.Println("Output:", myAtoi("  -0000000000012345678"))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了82.68%的用户
//内存消耗 :2.3 MB, 在所有 Go 提交中击败了47.59%的用户
func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 || (str[0] != '+' && str[0] != '-' && !unicode.IsDigit(rune(str[0]))) {
		return 0
	}
	isMinus := false
	start, end := 0, len(str)
	if str[0] == '-' {
		isMinus = true
		start = 1
	} else if str[0] == '+' {
		start = 1
	}
	for i := start; i < len(str); i++ {
		if !unicode.IsDigit(rune(str[i])) {
			end = i
			break
		}
	}
	return parseInt(str[start:end], isMinus)
}

func parseInt(str string, isMinus bool) int {
	var integer int
	maxint := "2147483647"
	minint := "2147483648"
	str = strings.TrimLeft(str, "0")
	if len(str) > 10 {
		if isMinus {
			return math.MinInt32
		}
		return math.MaxInt32
	} else if len(str) == 10 {
		if isMinus && str > minint {
			return math.MinInt32
		} else if !isMinus && str > maxint {
			return math.MaxInt32
		}
	}
	digit := 1
	for i := len(str) - 1; i >= 0; i-- {
		integer += int(str[i]-'0') * digit
		digit *= 10
	}
	if isMinus {
		integer *= -1
	}
	return integer
}
