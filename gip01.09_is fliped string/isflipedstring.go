// https://leetcode-cn.com/problems/string-rotation-lcci/
package main

import "fmt"

func main() {
	s1, s2 := "waterbottle", "erbottlewar"
	fmt.Println("Output: ", isFlipedString(s1, s2))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了27.39%的用户
//内存消耗 :6.2 MB, 在所有 Go 提交中击败了100.00%的用户
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	} else if s1 == s2 {
		return true
	}

	for i := range s2 {
		if s2[i] == s1[0] {
			if s1 == s2[i:]+s2[0:i] {
				return true
			}
		}
	}
	return false
}
