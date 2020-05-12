// https://leetcode-cn.com/problems/one-away-lcci/submissions/
package main

import "fmt"

func main() {
	first, second := "teacher", "beacher"
	fmt.Println("Output: ", oneEditAway(first, second))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了64.46%的用户
//内存消耗 :2.5 MB, 在所有 Go 提交中击败了100.00%的用户
func oneEditAway(first string, second string) bool {
	if len(first) == len(second) {
		rep := false
		for i := range first {
			if first[i] != second[i] {
				if rep {
					return false
				}
				rep = true
			}
		}
		return true
	} else if len(first) < len(second) {
		first, second = second, first
	}

	if len(first)-len(second) > 1 {
		return false
	} else if len(first) == 1 {
		return true
	}

	j := 0
	del := false
	for i := 0; i < len(first) && j < len(second); i++ {

		if first[i] != second[j] {
			if del {
				return false
			}
			del = true
		} else {
			j++
		}
	}

	return true
}
