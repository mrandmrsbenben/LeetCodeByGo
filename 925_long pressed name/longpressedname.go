package main

import "fmt"

func main() {
	// name := "ppyplrz"
	// typed := "pyypllrz"
	// name := "alex"
	// typed := "aaleex"
	// name := "saeed"
	// typed := "ssaaedd"
	// name := "leelee"
	// typed := "lleeelee"
	// name := "dfuyalc"
	// typed := "fuuyallc"
	name := "yyxbtsrs"
	typed := "yyyyxbbtssrs"
	fmt.Printf("Output: %v\n", isLongPressedName(name, typed))
}

func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	} else if name == "" || typed == "" {
		return false
	} else if name[0] != typed[0] || name[len(name)-1] != typed[len(typed)-1] {
		return false
	}
	i, j := 1, 1
	for i < len(typed) && j < len(name) {
		if name[j] == typed[i] {
			i++
			j++
		} else {
			if name[j-1] == typed[i] {
				i++
			} else {
				return false
			}
		}
	}
	if i == len(typed) && j != len(name) {
		return false
	}
	return true
}
