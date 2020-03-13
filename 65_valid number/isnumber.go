package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ".2e81" //".1" "-.1" "1." " 005047e+6""46.e3" ".2e81"
	fmt.Println("Output: ", isNumber(s))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.3 MB, 在所有 Go 提交中击败了56.52%的用户
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	switch strings.Count(s, "e") {
	case 0:
		return isDecimal(s)
	case 1:
		strs := strings.Split(s, "e")
		return isDecimal(strs[0]) && isInteger(strs[1])
	default:
		return false
	}
}

func isDecimal(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		if len(s) == 1 {
			return false
		}
		s = s[1:]
	}
	if len(s) == 1 && s[0] == '.' {
		return false
	}
	point := false
	for i := range s {
		if s[i] == '.' && point {
			return false
		} else if s[i] == '.' {
			point = true
		} else if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func isInteger(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		if len(s) == 1 {
			return false
		}
		s = s[1:]
	}
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
