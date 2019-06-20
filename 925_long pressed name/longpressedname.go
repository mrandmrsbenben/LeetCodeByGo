package main

import "fmt"

func main() {
	name := "pyplrz"
	typed := "ppyypllr"
	// name := "alex"
	// typed := "aaleex"
	// name := "saeed"
	// typed := "ssaaedd"
	// name := "leelee"
	// typed := "lleeelee"
	// name := "dfuyalc"
	// typed := "fuuyallc"
	// name := "yyxbtsrs"
	// typed := "yyyyxbbtssrs"
	fmt.Printf("Output: %v\n", isLongPressedName(name, typed))
}

func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	} else if name == "" || typed == "" {
		return false
	}
	for i, j := 0, 0; i < len(typed) && j < len(name); i++ {
		if name[j] != typed[i] {
			if i == 0 {
				return false
			}
			j++
		} else if j+1 < len(name) && name[j] == name[j+1] {
			j++
		}
		if i == len(typed)-1 && j != len(name)-1 {
			return false
		} else if name[j] != typed[i] {
			return false
		} else if j == len(name)-1 {
			break
		}
	}
	return true
}
