package main

import "fmt"

func main() {
	s := "RLRRLLRLRL"
	// s := "RLLLLRRRLR"
	// s := "LLLLRRRR"
	// s := "RRLRRLRLLLRL"
	fmt.Println("Output: ", balancedStringSplit(s))
}

//执行用时 :0 ms, 在所有 golang 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 golang 提交中击败了100.00%的用户
func balancedStringSplit(s string) int {
	ret := 0
	cnt := 0
	for _, c := range s {
		if c == 'R' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			ret++
		}
	}
	return ret
}
