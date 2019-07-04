package main

import "fmt"

func main() {
	fmt.Println("Output:", minAddToMakeValid(""))
	fmt.Println("Output:", minAddToMakeValid("()"))
	fmt.Println("Output:", minAddToMakeValid("())"))
	fmt.Println("Output:", minAddToMakeValid(")))"))
	fmt.Println("Output:", minAddToMakeValid("((("))
	fmt.Println("Output:", minAddToMakeValid("()))(("))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 Go 提交中击败了50.00%的用户
func minAddToMakeValid(S string) int {
	var l, r int
	for i := range S {
		if S[i] == '(' {
			l++
		} else {
			if l > 0 {
				l--
			} else {
				r++
			}
		}
	}
	return l + r
}
