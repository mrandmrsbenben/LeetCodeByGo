package main

import "fmt"

// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
func main() {
	fmt.Println("Output: ", lengthOfLongestSubstring("abccbcbb"))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :3.2 MB, 在所有 Go 提交中击败了100.00%的用户
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 1
	dic := make(map[rune]int)
	i, j := 0, 0
	var c rune
	for j, c = range s {
		if last, ok := dic[c]; ok {
			if j-i > max {
				max = j - i
			}
			if last+1 > i {
				i = last + 1
			}
		}
		dic[c] = j
	}
	if j-i+1 > max {
		max = j - i + 1
	}
	return max
}
