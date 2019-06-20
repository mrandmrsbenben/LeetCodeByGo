package main

import (
	"fmt"
)

func main() {
	S := "bxj##tw"
	T := "bxo#j##tw"
	fmt.Printf("Output: %v\n", backspaceCompare(S, T))
}

func backspaceCompare(S string, T string) bool {
	backStr := func(str string) string {
		back := ""
		bcnt := 0
		for i := len(str); i > 0; i-- {
			if str[i-1:i] == "#" {
				bcnt = bcnt + 1
			} else {
				if bcnt == 0 {
					back = str[i-1:i] + back
				} else {
					bcnt = bcnt - 1
				}
			}
		}
		return back
	}
	return backStr(S) == backStr(T)
}
